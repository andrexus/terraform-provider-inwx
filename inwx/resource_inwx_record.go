package inwx

import (
	"fmt"
	"log"

	"strconv"

	"github.com/andrexus/goinwx"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceINWXRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceINWXRecordCreate,
		Read:   resourceINWXRecordRead,
		Update: resourceINWXRecordUpdate,
		Delete: resourceINWXRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
			},

			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceINWXRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goinwx.Client)

	// Create the new record
	newRecord := &goinwx.NameserverRecordRequest{
		Domain:  d.Get("domain").(string),
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Content: d.Get("value").(string),
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newRecord.Ttl = ttl.(int)
	}

	log.Printf("[DEBUG] INWX Record create configuration: %#v", newRecord)

	recId, err := client.Nameservers.CreateRecord(newRecord)

	if err != nil {
		return fmt.Errorf("Failed to create INWX Record: %s", err)
	}

	d.SetId(strconv.Itoa(recId))
	log.Printf("[INFO] record ID: %s", d.Id())

	return resourceINWXRecordRead(d, meta)
}

func resourceINWXRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goinwx.Client)

	var domain *goinwx.NameserverDomain
	var rec *goinwx.NameserverRecord

	listResp, err := client.Nameservers.List("")
	if err != nil {
		return fmt.Errorf("Couldn't get list of INWX nameserver domains: %s", err)
	}

	for _, domainItem := range listResp.Domains {
		resp, err := client.Nameservers.Info("", domainItem.RoId)
		if err != nil {
			return fmt.Errorf("Couldn't get INWX nameserver domain info: %s", err)
		}

		for _, record := range resp.Records {
			if strconv.Itoa(record.Id) == d.Id() {
				rec = &record
				domain = &domainItem
				break
			}
		}
	}

	if rec == nil {
		return fmt.Errorf("Couldn't find INWX Record: %s", d.Id())
	}

	d.Set("domain", domain.Domain)
	d.Set("name", rec.Name)
	d.Set("type", rec.Type)
	d.Set("value", rec.Content)
	d.Set("ttl", rec.Ttl)
	d.Set("priority", rec.Prio)

	return nil
}

func resourceINWXRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goinwx.Client)

	recId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Can't convert recordId to int: %s", d.Id())
	}

	updateRecord := &goinwx.NameserverRecordRequest{}

	if attr, ok := d.GetOk("name"); ok {
		updateRecord.Name = attr.(string)
	}

	if attr, ok := d.GetOk("type"); ok {
		updateRecord.Type = attr.(string)
	}

	if attr, ok := d.GetOk("value"); ok {
		updateRecord.Content = attr.(string)
	}

	if attr, ok := d.GetOk("ttl"); ok {
		updateRecord.Ttl = attr.(int)
	}

	log.Printf("[DEBUG] INWX Record update configuration: %#v", updateRecord)

	err = client.Nameservers.UpdateRecord(recId, updateRecord)
	if err != nil {
		return fmt.Errorf("Failed to update INWX Record: %s", err)
	}

	return resourceINWXRecordRead(d, meta)
}

func resourceINWXRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goinwx.Client)

	log.Printf("[INFO] Deleting INWX Record: %s, %s", d.Get("domain").(string), d.Id())

	recId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Can't convert recordId to int: %s", d.Id())
	}
	err = client.Nameservers.DeleteRecord(recId)

	if err != nil {
		return fmt.Errorf("Error deleting INWX Record: %s", err)
	}

	return nil
}

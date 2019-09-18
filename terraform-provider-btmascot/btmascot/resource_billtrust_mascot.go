package btmascot

import (
	"fmt"
	"log"
	"strings"

	"github.com/billtrust/meetup-terraform-provider/api-client/models"

	billtrustclient "github.com/billtrust/meetup-terraform-provider/api-client/client"
	"github.com/billtrust/meetup-terraform-provider/api-client/client/operations"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMascot() *schema.Resource {
	return &schema.Resource{
		Create: mascotCreate,
		Update: mascotUpdate,
		Delete: mascotDelete,
		Read:   mascotRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// Validate will run against your TF before making calls to the API. It can have any logic in here
				//and if there's anything in the `errs` array, TF will fail
				ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
					if !strings.Contains(v.(string), "bill") {
						errs = append(errs, fmt.Errorf("%q must contain \"bill\"", k))
					}
					return
				},
			},
		},
	}
}

func mascotCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*billtrustclient.BilltrustMascot)

	// These varaibles are retrieved from the resource described in out .tf file
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	createMascotParams := operations.NewCreateMascotParams()
	createMascotParams.CreateRequest = &models.CreateMascotRequest{}
	createMascotParams.CreateRequest.Name = name
	createMascotParams.CreateRequest.Description = description

	mascot, err := client.Operations.CreateMascot(createMascotParams)

	if err != nil {
		return err
	}

	log.Println("=================== PRINTING CREATED MASCOT")
	log.Println(mascot)

	d.SetId(mascot.Payload.ID)

	return nil
}

func mascotUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*billtrustclient.BilltrustMascot)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mascotID := d.Id()

	updateMascotParams := operations.NewUpdateMascotParams()
	updateMascotParams.UpdateRequest = &models.UpdateMascotRequest{}
	updateMascotParams.UpdateRequest.Name = name
	updateMascotParams.UpdateRequest.Description = description
	updateMascotParams.ID = mascotID

	nil, err := client.Operations.UpdateMascot(updateMascotParams)

	if err != nil {
		return err
	}

	return nil
}

func mascotDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*billtrustclient.BilltrustMascot)

	// Since the resource exists in the state, we can get the ID from there to use it in the API calls against
	//the service
	mascotID := d.Id()

	deleteMascotParams := operations.NewDeleteMascotParams()
	deleteMascotParams.ID = mascotID

	nil, err := client.Operations.DeleteMascot(deleteMascotParams)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func mascotRead(d *schema.ResourceData, m interface{}) error {
	log.Println("================ READING MASCOT")
	client := m.(*billtrustclient.BilltrustMascot)

	mascotID := d.Id()

	readMascotRequest := operations.NewGetMascotParams()
	readMascotRequest.ID = mascotID

	mascot, err := client.Operations.GetMascot(readMascotRequest)

	if err != nil {
		// If we return an error, Terraform will fail and the state will not be updated. In this case, since we know
		// the API and the only error scenario is a 404, we are ok with a catchall. Setting the id to "" makes TF
		// removed the resource from the state and tries to resync
		d.SetId("")
		return nil
	}

	log.Println(mascot)

	return nil
}

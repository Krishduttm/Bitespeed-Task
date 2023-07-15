package routes

import (
	"bitespeedtask/config"
	"bitespeedtask/models"
	"bitespeedtask/resources"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	router := route.Group("/bitespeed")
	{
		router.GET("/getAll", func(c *gin.Context) {
			var contacts []models.Contact

			res := config.DB.Find(&contacts)
			if res.Error != nil {
				fmt.Println("Error while retreiving contacts")
			}

			c.JSON(http.StatusOK, gin.H{"data": contacts})
		})
	}

	{
		router.POST("/identify", func(c *gin.Context) {

			var input resources.GetContactDetails
			var primaryContact models.Contact
			var secondaryContact models.Contact
			var contacts []models.Contact
			flag := false

			err := c.BindJSON(&input)
			if err != nil {
				panic("error while binding")
			}

			query := config.DB.Debug().Where("(phone_number = ? OR email = ?) AND link_precedence = ?", input.PhoneNumber, input.Email, "secondary").Find(&secondaryContact)
			fmt.Println("query: ", secondaryContact)
			if query.Error != nil {
				fmt.Println("Error while retreiving primary contacts")
			}

			linkedId := secondaryContact.LinkedIn
			var contact models.Contact

			if secondaryContact.LinkedIn == 0 {
				query := config.DB.Debug().Where("(phone_number = ? OR email = ?) AND link_precedence = ?", input.PhoneNumber, input.Email, "primary").Find(&primaryContact)
				fmt.Println("secondary bt primary: ", primaryContact)
				linkedId = primaryContact.Id
				if query.Error != nil {
					fmt.Println("Error while retreiving primary contacts")
				}
			}

			if primaryContact.Id == 0 && secondaryContact.Id == 0 {
				create := config.DB.Debug().Create(
					&models.Contact{
						Email:          input.Email,
						PhoneNumber:    input.PhoneNumber,
						LinkPrecedence: "primary",
						CreatedAt:      time.Now(),
					})
				flag = true
				fmt.Println("query: ", &models.Contact{
					Email:          input.Email,
					PhoneNumber:    input.PhoneNumber,
					LinkPrecedence: "primary",
					CreatedAt:      time.Now(),
				})
				if create.Error != nil {
					fmt.Println("Error while creating primary contacts")
				}
			}

			if input.Email != "" && input.PhoneNumber != "" && !flag {
				if config.DB.Debug().Model(&contact).Where("email = ? AND phone_number = ? AND link_precedence = ?", input.Email, input.PhoneNumber, "secondary").Updates(&contact).RowsAffected == 0 {
					config.DB.Debug().Create(&models.Contact{
						Email:          input.Email,
						PhoneNumber:    input.PhoneNumber,
						LinkedIn:       linkedId,
						LinkPrecedence: "secondary",
					})
				}
			}

			query2 := config.DB.Debug().Where("id = ?", linkedId).Find(&primaryContact)
			if query2.Error != nil {
				fmt.Println("Error while retreiving scondary contacts")
			}

			query3 := config.DB.Debug().Where("linked_in = ?", linkedId).Find(&contacts)
			if query3.Error != nil {
				fmt.Println("Error while retreiving all scondary contacts")
			}
			output := resources.ContactDetailsOutput{
				PrimaryContactID: linkedId,
			}

			output.Emails = append(output.Emails, primaryContact.Email)
			output.PhoneNumbers = append(output.PhoneNumbers, primaryContact.PhoneNumber)
			for _, value := range contacts {
				output.Emails = append(output.Emails, value.Email)
				if output.PhoneNumbers[0] != value.PhoneNumber {
					output.PhoneNumbers = append(output.PhoneNumbers, value.PhoneNumber)
				}
				output.SecondaryContactIds = append(output.SecondaryContactIds, value.Id)
			}
			c.JSON(http.StatusOK, gin.H{"data": output})
		})
	}
}

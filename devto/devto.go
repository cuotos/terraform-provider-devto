package devto

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cuotos/devto-go/devto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceArticle() *schema.Resource {
	resource := &schema.Resource{}

	resource.CreateContext = resourceArticleCreate
	resource.ReadContext = resourceArticleRead
	resource.UpdateContext = resourceArticleUpdate
	resource.DeleteContext = resourceArticleDelete

	resource.Schema = map[string]*schema.Schema{
		"title": {
			Type:     schema.TypeString,
			Required: true,
		},
		"markdown_body": {
			Type:     schema.TypeString,
			Required: true,
		},
		"published": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}

	return resource
}

func resourceArticleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*devto.Client)

	title := data.Get("title").(string)
	bodyMarkdown := data.Get("markdown_body").(string)
	published := data.Get("published").(bool)

	article := devto.CreateArticle{
		Title:        title,
		BodyMarkdown: bodyMarkdown,
		Published:    published,
	}

	createdArticle, _, err := client.CreateArticle(article)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(fmt.Sprintf("%d", createdArticle.ID))

	return nil
}

func resourceArticleRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*devto.Client)

	id, err := strconv.Atoi(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	article, found, _, err := client.GetUserArticleByID(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if found {
		data.SetId(data.Id())
		data.Set("title", article.Title)
	}

	return nil

}

func resourceArticleUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*devto.Client)

	title := data.Get("title").(string)
	bodyMarkdown := data.Get("markdown_body").(string)

	article := devto.CreateArticle{
		Title:        title,
		BodyMarkdown: bodyMarkdown,
	}

	id, err := strconv.Atoi(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	createdArticle, _, err := client.UpdateArticle(id, article)
	if err != nil {
		diag.FromErr(err)
	}

	data.SetId(fmt.Sprintf("%d", createdArticle.ID))

	return nil
}

func resourceArticleDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {

	// TODO
	return diag.Errorf("TODO: Delete article not implemented yet")
}

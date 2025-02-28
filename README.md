# Meilisearch API Example

This is an example project demonstrating how to interact with Meilisearch via a REST API. It includes basic functionality to search, add, delete, and update movies stored in Meilisearch.

## Requirements

Before running this project, you need to have the following installed:

- Go (1.18+)
- Docker

## How to run project

Run local meilisearch with Docker:

```bash
docker run -d --name meilisearch \
  -p 7700:7700 \
  -e MEILI_MASTER_KEY='my_master_key' \
  getmeili/meilisearch:v1.8
```

Run project

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/26016872-d8812d14-77df-44e7-aef0-c6e86d1a2f51?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D26016872-d8812d14-77df-44e7-aef0-c6e86d1a2f51%26entityType%3Dcollection%26workspaceId%3D2f6a042b-41bf-4696-bf48-afdd8de1ba58)
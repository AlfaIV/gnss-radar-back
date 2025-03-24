import json
import os
# from pprint import pprint
# import jsonref
from app.core.config import configs


def gen_doc(open_api_path):
    # Загрузка OpenAPI спецификации
    client = OpenAPIClient.from_file(open_api_path)

    # Генерация Markdown (примерный код, зависит от реализации)
    markdown_docs = client.generate_markdown()

    # Сохранение в файл
    with open('api_documentation.md', 'w', encoding='utf-8') as file:
        file.write(markdown_docs)

    print("Markdown документация успешно создана!")


def resolve_refs(schema):
    """Рекурсивно разрешает $ref в схеме."""
    if isinstance(schema, dict):
        if "$ref" in schema:
            # Разрешаем ссылку
            ref = schema["$ref"]
            return resolve_refs(jsonref.replace_refs(schema))
        else:
            # Рекурсивно обрабатываем вложенные элементы
            return {k: resolve_refs(v) for k, v in schema.items()}
    elif isinstance(schema, list):
        # Рекурсивно обрабатываем элементы списка
        return [resolve_refs(item) for item in schema]
    else:
        return schema

def generate_markdown_docs(openapi):
    """Генерация Markdown документации."""
    openapi_schema = openapi
    markdown = []

    # Разрешаем все $ref в схеме
    resolved_schema = resolve_refs(openapi_schema)

    # Заголовок API
    markdown.append(f"# {resolved_schema['info']['title']}\n")
    markdown.append(f"**Version:** {resolved_schema['info']['version']}\n")
    markdown.append(f"**Description:** {resolved_schema['info'].get('description', '')}\n\n")

    # Эндпоинты
    for path, path_item in resolved_schema['paths'].items():
        markdown.append(f"## `{path}`\n")
        for method, operation in path_item.items():
            if method not in ['get', 'post', 'put', 'delete', 'patch']:
                continue

            markdown.append(f"### `{method.upper()}`\n")
            markdown.append(f"**Summary:** {operation.get('summary', '')}\n")
            markdown.append(f"**Description:** {operation.get('description', '')}\n\n")

            # Параметры
            if 'parameters' in operation:
                markdown.append("#### Parameters\n")
                for param in operation['parameters']:
                    markdown.append(f"- **{param['name']}** (in: {param['in']}, type: {param.get('schema', {}).get('type', 'unknown')}): {param.get('description', '')}\n")
                markdown.append("\n")

            # Тело запроса
            if 'requestBody' in operation:
                markdown.append("#### Request Body\n")
                for content_type, content in operation['requestBody']['content'].items():
                    markdown.append(f"- **Content Type:** `{content_type}`\n")
                    if 'schema' in content:
                        schema = content['schema']
                        markdown.append(f"  **Schema:**\n")
                        markdown.append(f"  ```json\n{json.dumps(schema, indent=2)}\n  ```\n")
                markdown.append("\n")

            # Ответы
            if 'responses' in operation:
                markdown.append("#### Responses\n")
                for status_code, response in operation['responses'].items():
                    markdown.append(f"- **{status_code}**: {response.get('description', '')}\n")
                    if 'content' in response:
                        for content_type, content in response['content'].items():
                            markdown.append(f"  **Content Type:** `{content_type}`\n")
                            if 'schema' in content:
                                schema = content['schema']
                                markdown.append(f"  **Schema:**\n")
                                markdown.append(f"  ```json\n{json.dumps(schema, indent=2)}\n  ```\n")
                markdown.append("\n")

    # Сохранение в файл
    doc_path = os.path.join(configs.PROJECT_ROOT, "docs")
    os.makedirs(doc_path, exist_ok=True)  # Создаем папку, если её нет
    with open(f"{doc_path}/api_documentation.md", "w", encoding="utf-8") as file:
        file.write("\n".join(markdown))

    print("Markdown документация успешно создана!")

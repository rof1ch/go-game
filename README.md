# Go-Game 🎮

Небольшая консольная игра-ходилка

### Конфигурации

```json
[
	{
		"name": "Дом", // Название локации
		"description": "Мой дом", // Описание локации
		"is_open": true, // Открыта ли локация изначально
		"zones": [
			// Зоны внутри локации
			{
				"name": "Кухня", // Название зоны
				"items": [
					// Обьекты внутри зоны
					{
						"type": "key", // Тип обьекта
						"name": "Ключ от сада", // Название обьекта
						"location_name": "Сад" // Для типа key необходим location_name
					},
					{
						"type": "health_potion",
						"name": "Зелье здоровья",
						"health": 20 // Для зелья здоровья необходимое поле
					}
				]
			},
			{
				"name": "Гостиная", // Название зоны
				"items": [
					{
						"type": "weapon",
						"name": "Нож",
						"damage": 10 // Для оружия необходимое поле урона
					}
				]
			}
		],
		"locations": [
			// На какие локации можно уйти с этой
			{
				"name": "Сад" // Необходимо только имя локации
			}
		]
	}
]
```

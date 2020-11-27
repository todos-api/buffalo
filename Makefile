api = :3000/todos

run: run1 run2

run1:
	# create a to-do list, and add items to it, then delete it
	$(eval list_id := \
		$(shell http --body POST $(api) title="List 4" created_by="user1" | \
		jq --raw-output .id))
	http POST $(api)/$(list_id)/items name="Name 1"
	http POST $(api)/$(list_id)/items name="Name 2"
	http $(api)/$(list_id)
	http $(api)/$(list_id)/items
	http --verbose DELETE $(api)/$(list_id) < /dev/null

run2:
	# create a bunch of to-do lists, and then delete them
	http $(api)
	http POST $(api) title="List 1" created_by="user1"
	http POST $(api) title="List 2" created_by="user1"
	http POST $(api) title="List 3" created_by="user2"
	http $(api)
	http --body GET $(api) | jq --raw-output .[].id | while read ID; do \
		http --verbose DELETE $(api)/$${ID} < /dev/null; \
	done
	http $(api)

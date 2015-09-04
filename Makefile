
post:
	go run tasks/make_post_task.go

serve:
	goapp serve --host 0.0.0.0 app/

list-posts:
	curl localhost:8080/posts/list

list-posts-pretty:
	curl localhost:8080/posts/list | prettyjson

create-post:
	curl -XPOST localhost:8080/posts/create --data "car_plate=AAA-1233&message=HAte you" | prettyjson
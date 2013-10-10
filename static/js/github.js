(function() {
	var templates = {
		"PushEvent": Handlebars.compile($("#github-PushEvent-template").html()),
		"CreateEvent": Handlebars.compile($("#github-CreateEvent-template").html()),
		"DeleteEvent": Handlebars.compile($("#github-DeleteEvent-template").html()),
		"IssuesEvent": Handlebars.compile($("#github-IssuesEvent-template").html()),
		"IssueCommentEvent": Handlebars.compile($("#github-IssueCommentEvent-template").html()),
	};

	var userTemplate = Handlebars.compile($("#github-user-template").html());

	AboutMe.github = function() {
		$.get("/me/github/events", function(data) {
			$.each(data, function(index, item) {
				if (templates[item.type]) {
					$('#github-cards').append(AboutMe.cardTemplate({ template: templates[item.type](item) }));
				} else {
					// report that I don't support this type of event.
					//$.post("/me/unsupported", item, "json");
				}
			});
		}, "json");
		$.get("/me/github/me", function(data) {
			$('#github-user').html(userTemplate(data));
		}, "json");
	}
})();
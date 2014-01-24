(function() {
	var templates = {
		"PushEvent": Handlebars.compile($("#github-PushEvent-template").html()),
		"CreateEvent": Handlebars.compile($("#github-CreateEvent-template").html()),
		"DeleteEvent": Handlebars.compile($("#github-DeleteEvent-template").html()),
		"IssuesEvent": Handlebars.compile($("#github-IssuesEvent-template").html()),
		"IssueCommentEvent": Handlebars.compile($("#github-IssueCommentEvent-template").html()),
		"FollowEvent": Handlebars.compile($("#github-FollowEvent-template").html()),
		"PullRequestEvent": Handlebars.compile($("#github-PullRequestEvent-template").html()),
		"MemberEvent": Handlebars.compile($("#github-MemberEvent-template").html()),
		"PullRequestReviewCommentEvent": Handlebars.compile($("#github-PullRequestReviewCommentEvent-template").html()),
		"WatchEvent": Handlebars.compile($("#github-WatchEvent-template").html()),
		"PublicEvent": Handlebars.compile($("#github-PublicEvent-template").html()),
		"CommitCommentEvent": Handlebars.compile($("#github-CommitCommentEvent-template").html()),
		"GollumEvent": Handlebars.compile($("#github-GollumEvent-template").html())
	};

	var userTemplate = Handlebars.compile($("#github-user-template").html());

	AboutMe.github = function() {
		$('#github-link').html('<a href="https://www.github.com/' + AboutMe.config.github + '" target="_blank">github</a>');
		$.get("/github/events", function(data) {
			_.each(data, function(item) {
				if (templates[item.type]) {
					var githubItem = {
						template: $(AboutMe.cardTemplate({ template: templates[item.type](item) })),
						timestamp: new Date(item.created_at).getTime()
					};
					AboutMe.events.github.push(githubItem);
					AboutMe.events.tech.push(githubItem);
					AboutMe.events.all.push(githubItem);
					AboutMe.getElementToAdd("github", "tech").append(githubItem.template);
				} else {
					// report that I don't support this type of event.
					$.post("/unsupported", JSON.stringify(item));
				}
			});
			AboutMe.events.tech = AboutMe.sortEvents(AboutMe.events.tech);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
		$.get("/github/me", function(data) {
			$('#github-user').html(userTemplate(data));
		}, "json");
	}
})();
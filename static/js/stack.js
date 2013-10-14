(function() {
	var templates = {
		"commented": Handlebars.compile($("#stack-commented-template").html()),
		"answered": Handlebars.compile($("#stack-answered-template").html()),
		"badge": Handlebars.compile($("#stack-badge-template").html()),
		"suggested": Handlebars.compile($("#stack-suggested-template").html()),
	};
	var userTemplate = Handlebars.compile($("#stack-user-template").html());

	AboutMe.stack = function() {
		$('#stack-link').html('<a href="http://stackoverflow.com/users/' + AboutMe.config.stack + '" target="_blank">stackoverflow</a>');
		$.get("/me/stack/me", function(data) {
			_.each(data.items, function(item) {
				$('#stack-user').html(userTemplate(item));
			});
		}, "json");
		$.get("/me/stack/timeline", function(data) {
			_.each(data.items, function(item) {
				if (templates[item.timeline_type]) {
					var stackItem = {
						template: $(AboutMe.cardTemplate({template: templates[item.timeline_type](item)})),
						timestamp: item.creation_date * 1000
					};
					AboutMe.events.stack.push(stackItem);
					AboutMe.events.tech.push(stackItem);
					AboutMe.events.all.push(stackItem);
					AboutMe.getElementToAdd("stack", "tech").append(stackItem.template);
				} else {
					$.post("/me/unsupported", JSON.stringify(item));
				}
			});
			AboutMe.events.tech = AboutMe.sortEvents(AboutMe.events.tech);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
	}
})();
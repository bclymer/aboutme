(function() {
	var userTemplate = Handlebars.compile($("#stack-user-template").html());
	var timelineTemplate = Handlebars.compile($("#stack-timeline-template").html());

	AboutMe.stack = function() {
		$('#stack-link').html('<a href="http://stackoverflow.com/users/' + AboutMe.config.stack + '" target="_blank">stackoverflow</a>');
		$.get("/me/stack/me", function(data) {
			_.each(data.items, function(item) {
				$('#stack-user').html(userTemplate(item));
			});
		}, "json");
		$.get("/me/stack/timeline", function(data) {
			_.each(data.items, function(item) {
				var stackItem = {
					template: $(AboutMe.cardTemplate({template: timelineTemplate(item)})),
					timestamp: item.creation_date * 1000
				};
				AboutMe.events.stack.push(stackItem);
				AboutMe.events.tech.push(stackItem);
				AboutMe.events.all.push(stackItem);
				$('#stack-cards').append(stackItem.template);
			});
			AboutMe.events.tech = AboutMe.sortEvents(AboutMe.events.tech);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
	}
})();
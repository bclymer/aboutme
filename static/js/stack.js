(function() {
	var userTemplate = Handlebars.compile($("#stack-user-template").html());
	var timelineTemplate = Handlebars.compile($("#stack-timeline-template").html());

	AboutMe.stack = function() {
		$.get("/me/stack/me", function(data) {
			$.each(data.items, function(index, item) {
				$('#stack-user').html(userTemplate(item));
			});
		}, "json");
		$.get("/me/stack/timeline", function(data) {
			$.each(data.items, function(index, item) {
				$('#stack-cards').append(AboutMe.cardTemplate({template: timelineTemplate(item)}));
			});
		}, "json");
	}
})();
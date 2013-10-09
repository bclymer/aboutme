(function() {
	var timelineTemplate = Handlebars.compile($("#stack-template").html());

	AboutMe.stack = function() {
		$.get("/me/stack", function(data) {
			$.each(data.items, function(index, item) {
				$('#stack-cards').append(timelineTemplate(item));
			});
		}, "json");
	}
})();
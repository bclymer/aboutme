(function() {
	var infoTemplate = Handlebars.compile($("#me-info-template").html());

	AboutMe.me = function() {
		$.get("/me/info", function(data) {
			var fakeTime = 2000000000000;
			_.each(data, function(item) {
				var meItem = {
					template: $(AboutMe.cardTemplate({template: infoTemplate(item)})),
					timestamp: fakeTime--
				};
				AboutMe.events.me.push(meItem);
				AboutMe.events.social.push(meItem);
				AboutMe.events.all.push(meItem);
				AboutMe.getElementToAdd("me", "social").append(meItem.template);
			});
			AboutMe.events.social = AboutMe.sortEvents(AboutMe.events.social);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
	}
})();
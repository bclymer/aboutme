(function() {
	var userTemplate = Handlebars.compile($("#facebook-user-template").html());
	var timelineTemplate = Handlebars.compile($("#facebook-feed-template").html());

	AboutMe.facebook = function() {
		$('#facebook-link').html('<a href="https://www.facebook.com/' + AboutMe.config.facebook + '" target="_blank">facebook</a>');
		return;
		$.get("/me/facebook/feed", function(data) {
			_.each(data.data, function(item) {
				var facebookItem = {
					template: $(AboutMe.cardTemplate({template: timelineTemplate(item)})),
					timestamp: new Date(item.created_time).getTime()
				};
				AboutMe.events.facebook.push(facebookItem);
				AboutMe.events.social.push(facebookItem);
				AboutMe.events.all.push(facebookItem);
				AboutMe.getElementToAdd("facebook", "social").append(facebookItem.template);
			});
			AboutMe.events.social = AboutMe.sortEvents(AboutMe.events.social);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
	}
})();
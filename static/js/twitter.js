(function() {
	var userTemplate = Handlebars.compile($("#twitter-user-template").html());
	var timelineTemplate = Handlebars.compile($("#twitter-timeline-template").html());

	AboutMe.twitter = function() {
		$('#twitter-link').html('<a href="https://twitter.com/' + AboutMe.config.twitter + '" target="_blank">twitter</a>');
		$.get("/me/twitter/timeline", function(data) {
			if (data.length > 0) {
				$('#twitter-user').html(userTemplate(data[0]));
			}
			_.each(data, function(item) {
				var twitterItem = {
					template: $(AboutMe.cardTemplate({template: timelineTemplate(item)})),
					timestamp: new Date(item.created_at).getTime()
				};
				AboutMe.events.stack.push(twitterItem);
				AboutMe.events.social.push(twitterItem);
				AboutMe.events.all.push(twitterItem);
				$('#twitter-cards').append(twitterItem.template);
			});
			AboutMe.events.social = AboutMe.sortEvents(AboutMe.events.social);
			AboutMe.events.all = AboutMe.sortEvents(AboutMe.events.all);
		}, "json");
	}
})();
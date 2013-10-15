(function() {
	
	enquire.register("(min-width: 1600px)", {
		match : function() {
			if (AboutMe.responsiveState >= 3) {
				moveListToElement(AboutMe.events.stack, $('#stack-cards'));
				moveListToElement(AboutMe.events.github, $('#github-cards'));
			}
			if (AboutMe.responsiveState >= 2) {
				moveListToElement(AboutMe.events.twitter, $('#twitter-cards'));
				moveListToElement(AboutMe.events.facebook, $('#facebook-cards'));
			}
			AboutMe.responsiveState = 1;
		}
	});

	enquire.register("(min-width: 1200px) and (max-width: 1599px)", {
		match : function() {
			if (AboutMe.responsiveState == 1 || AboutMe.responsiveState == 4) {
				moveListToElement(AboutMe.events.social, $('#social-cards'));
			}
			if (AboutMe.responsiveState >= 3) {
				moveListToElement(AboutMe.events.stack, $('#stack-cards'));
				moveListToElement(AboutMe.events.github, $('#github-cards'));
			}
			AboutMe.responsiveState = 2;
		}
	});

	enquire.register("(min-width: 800px) and (max-width: 1199px)", {
		match : function() {
			if (AboutMe.responsiveState == 1 || AboutMe.responsiveState == 4) {
				moveListToElement(AboutMe.events.social, $('#social-cards'));
			}
			if (AboutMe.responsiveState <= 2 || AboutMe.responsiveState == 4) {
				moveListToElement(AboutMe.events.tech, $('#tech-cards'));
			}
			AboutMe.responsiveState = 3;
		}
	});

	enquire.register("(max-width: 799px)", {
		match : function() {
			moveListToElement(AboutMe.events.all, $('#all-cards'));
			AboutMe.responsiveState = 4;
		}
	});

	function moveListToElement(list, element) {
		_.each(list, function(event) {
			element.append(event.template);
		});
	}
	
})();
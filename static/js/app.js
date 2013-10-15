(function() {

	window.AboutMe = {};

	AboutMe.events = {
		stack: [],
		github: [],
		twitter: [],
		facebook: [],
		tech: [],
		social: [],
		all: [],
	};

	AboutMe.responsiveState = -1;

	AboutMe.sortEvents = function(events) {
		return _.sortBy(events, function(event) {
			return -event.timestamp;
		});
	}

	Handlebars.registerHelper('ifEq', function(v1, v2, options) {
		if(v1 === v2) {
			return options.fn(this);
		}
		return options.inverse(this);
	});

	Handlebars.registerHelper('timeagoHelper', function(time, options) {
		return options.fn({
			timeagoValue: AboutMe.timeago().timeDifference(time)
		});
	});

	Handlebars.registerHelper('plural', function(count, word, options) {
		return options.fn({
			word: (count == 1) ? word : word + "s"
		});
	});

	AboutMe.getElementToAdd = function(primary, secondary) {
		switch (AboutMe.responsiveState) {
			case 4:
				return $('#all-cards');
			case 3:
				return $('#' + secondary + '-cards');
			case 2:
				if (secondary == "tech") {
					return $('#' + primary + '-cards');
				} else {
					return $('#' + secondary + '-cards');
				}
			case 1:
				return $('#' + primary + '-cards');
		}
	};

	$(document).ready(function() {
		AboutMe.cardTemplate = Handlebars.compile($("#card-template").html());
		AboutMe.stack();
		AboutMe.github();
		AboutMe.twitter();
		AboutMe.facebook();
	});

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
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

	AboutMe.previousState = -1;

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

	$(document).ready(function() {
		AboutMe.cardTemplate = Handlebars.compile($("#card-template").html());
		AboutMe.stack();
		AboutMe.github();
		AboutMe.twitter();
		AboutMe.facebook();
	});

	enquire.register("(min-width: 1600px)", {
    	match : function() {
    		moveListToElement(AboutMe.events.stack, $('#stack-cards'));
    		moveListToElement(AboutMe.events.github, $('#github-cards'));
    		moveListToElement(AboutMe.events.twitter, $('#twitter-cards'));
    		moveListToElement(AboutMe.events.facebook, $('#facebook-cards'));
    	}
	});

	enquire.register("(min-width: 1200px) and (max-width: 1599px)", {
    	match : function() {
			moveListToElement(AboutMe.events.social, $('#social-cards'));
    		moveListToElement(AboutMe.events.stack, $('#stack-cards'));
    		moveListToElement(AboutMe.events.github, $('#github-cards'));
    	}
	});

	enquire.register("(min-width: 800px) and (max-width: 1199px)", {
    	match : function() {
			moveListToElement(AboutMe.events.social, $('#social-cards'));
    		moveListToElement(AboutMe.events.tech, $('#tech-cards'));
    	}
	});

	enquire.register("(max-width: 799px)", {
    	match : function() {
    		moveListToElement(AboutMe.events.all, $('#all-cards'));
    	}
	});

	function moveListToElement(list, element) {
		_.each(list, function(event) {
			element.append(event.template);
		});
	}

})();
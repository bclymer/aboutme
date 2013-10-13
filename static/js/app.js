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
/*
	enquire.register("(min-width: 1900px)", {
    	match : function() {
    		console.log("To 4 from " + AboutMe.previousState);
    		if (AboutMe.previousState <= 2) {
	    		moveListToElement(AboutMe.events.stack, $('#stack-cards'));
	    		moveListToElement(AboutMe.events.github, $('#github-cards'));
    		}
    		if (AboutMe.previousState <= 3) {
	    		moveListToElement(AboutMe.events.twitter, $('#twitter-cards'));
	    		moveListToElement(AboutMe.events.facebook, $('#facebook-cards'));
    		}
    		AboutMe.previousState = 4;
    	}
	});

	enquire.register("(min-width: 1800px) and (max-width: 1899px)", {
    	match : function() {
    		console.log("To 3 from " + AboutMe.previousState);
    		if (AboutMe.previousState >= 4) {
    			moveListToElement(AboutMe.events.social, $('#social-cards'));
    		} else if (AboutMe.previousState <= 2) {
	    		moveListToElement(AboutMe.events.stack, $('#stack-cards'));
	    		moveListToElement(AboutMe.events.github, $('#github-cards'));	
    		}
    		AboutMe.previousState = 3;
    	}
	});

	enquire.register("(min-width: 1700px) and (max-width: 1799px)", {
    	match : function() {
    		console.log("To 2 from " + AboutMe.previousState);
    		if (AboutMe.previousState >= 4) {
    			moveListToElement(AboutMe.events.social, $('#social-cards'));
    		}
    		if (AboutMe.previousState >= 3 || AboutMe.previousState <= 1) {
	    		moveListToElement(AboutMe.events.tech, $('#tech-cards'));
    		}
    		AboutMe.previousState = 2;
    	}
	});

	enquire.register("(max-width: 1699px)", {
    	match : function() {
    		console.log("To 1 from " + AboutMe.previousState);
    		if (AboutMe.previousState >= 2) {
    			moveListToElement(AboutMe.events.all, $('#all-cards'));
    		}
    		AboutMe.previousState = 1;
    	}
	});

	function moveListToElement(list, element) {
		_.each(list, function(event) {
			element.append(event.template);
		});
	}
*/
})();
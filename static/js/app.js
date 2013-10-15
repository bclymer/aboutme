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

})();
(function() {

	window.AboutMe = {};
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
	});

})();
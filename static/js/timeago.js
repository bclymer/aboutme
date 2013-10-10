/*
 * Timeago is a utility that converts unix dates to h
*/
(function() {
	AboutMe.timeago = function() {

		// modified http://stackoverflow.com/questions/6108819/javascript-timestamp-to-relative-time-eg-2-seconds-ago-one-week-ago-etc-best
		function timeDifference(previous) {

		    var secPerMinute = 60;
		    var secPerHour = secPerMinute * 60;
		    var secPerDay = secPerHour * 24;
		    var secPerMonth = secPerDay * 30;
		    var secPerYear = secPerDay * 365;
			var now = new Date();
			var utc = new Date(now.getTime() + now.getTimezoneOffset() * 60000).getTime();

			if (!Number(previous)) {
				previous = new Date(previous).getTime() / 1000;
			}
		    var elapsed = Math.round(utc / 1000) - previous;

		    if (elapsed < secPerMinute) {
		         return elapsed + ' seconds ago';   
		    }

		    else if (elapsed < secPerHour) {
		         return pluralize(Math.round(elapsed / secPerMinute), 'minute');
		    }

		    else if (elapsed < secPerDay ) {
		         return pluralize(Math.round(elapsed / secPerHour), 'hour');
		    }

		    else if (elapsed < secPerMonth) {
		        return pluralize(Math.round(elapsed / secPerDay), 'day');
		    }

		    else if (elapsed < secPerYear) {
		        return pluralize(Math.round(elapsed / secPerMonth), 'month');
		    }

		    else {
		        return pluralize(Math.round(elapsed / secPerYear ), 'year');
		    }
		}

		function pluralize(time, word) {
			return time + ' ' + (time > 1 ? word + 's' : word) + ' ago';   
		}

		return {
			timeDifference: timeDifference
		}
	}
})();
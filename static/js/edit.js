(function() {
	var infoTemplate = Handlebars.compile($('#me-info-template').html());

	$.get('/me/info', function(data) {
		_.each(data, function(item) {
			template = $(infoTemplate(item));
			$('#cards').append(template);
		});
	}, 'json');

	$('#save').click(function() {
		cards = [];
		_.each($('.card'), function(card) {
			var jCard = $(card);
			cardObject = {
				title: jCard.find('.title').val(),
				body: jCard.find('.body').val(),
				link: {
					url: jCard.find('.link-url').val(),
					text: jCard.find('.link-text').val(),
				}
			}
			cards.push(cardObject);
		});
		var auth = prompt('What\'s the auth though?');
		$.post('/edit/info?auth=' + auth, JSON.stringify(cards), 'json')
			.done(function(data) {
				alert(data);
			});
	});

	$('#add').click(function() {
		fakeObject = {
			title: 'Title',
			body: 'Body',
			link: {
				url: 'Link URL',
				text: 'Link Text',
			}
		}
		template = $(infoTemplate(fakeObject));
		$('#cards').append(template);
	});

	$(document).on('click', '.remove', function() {
		console.log($(this).parent().parent().remove()); // mmm this is nice
	});
})();
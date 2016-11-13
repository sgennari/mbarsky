var registration = $('#registration');
registration.hide();
// var login = $('#login');
// login.hide();

formOpen = false;
formButton = $('#start');
formButtonText = formButton.html();
formButton.click(function () {
	formOpen = !formOpen;
	if (formOpen) {
		// login.fadeIn();
		registration.fadeIn();
		registrationForm.trigger('reset');
		formButton.html('Close Registration');
	} else {
		$('.alert').fadeOut();
		// login.fadeOut();
		registration.fadeOut();
		formButton.html(formButtonText);
	}
})


var needsRefresh = false;
var registrationForm = $('#registration-form');
registrationForm.submit(function (event) {
	$('.alert').fadeOut();
	var url = '/api/v1/talks';
	$.ajax({
		type: 'POST',
		url: url,
		data: registrationForm.serialize(),
		success: function (data) {
			needsRefresh = true;
			registrationForm.append('<p class="alert alert-success"><strong>Well Done!</strong> You\'re registered!</p>');
			formButton.html('Register another talk?');
			return;
		},
		error: function () {
			registrationForm.append('<p class="alert alert-error"><strong>Oh Snap!</strong> You\'re missing a few fields! Try again!</p>');
		}
	});
	event.preventDefault();
});
// var loginForm = $('#login-form');

// loginForm.submit(function (event) {
// 	$.ajax();
// 	event.preventDefault();
// });

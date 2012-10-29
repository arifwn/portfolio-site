
$(".profile-pic img").popover(
	title: 'Contact QR Code',
	content: '<img src="/img/assets/contact_qr.png" width="250" alt="" />',
	placement: 'bottom',
).click((e) ->
	e.preventDefault()
)

$(".profile-pic img").tooltip(
	title: 'Click to toggle Contact QR Code',
	placement: 'right',
)
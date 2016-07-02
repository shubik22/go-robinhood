require('dotenv').config();

var fs = require('fs');
var Robinhood = require('robinhood');
var username = process.env.USERNAME;
var password = process.env.PASSWORD;

var endpoints = [
  'investment_profile',
  'accounts',
  'user',
  'dividends',
  'orders',
  'positions'
];

function downloadAndSave(robinhood, endpoint) {
	console.log('** Downloading ' + endpoint);

	robinhood[endpoint](function(err, response, body) {
		var filename = '../data/' + endpoint + '.json';
		var data = JSON.stringify(body);
		console.log('** Saving ' + endpoint);
		fs.writeFile(filename, data);
	});
}

var robinhood = Robinhood({
	username: username,
	password: password
}, function() {
	endpoints.forEach(downloadAndSave.bind(this, robinhood))
});

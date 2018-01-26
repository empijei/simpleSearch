//Disable enter to submit and add to clipboard instead
$(document).ready(function() {
	$('form input').keydown(function(event){
		//FIXME this does not prevent sending anymore
		if(event.keyCode == 13) {
			var results = document.getElementsByClassName("list-group-item");
			listClick({target:results[0].children[0]});
			event.preventDefault();
			return false;
		}
	});
});

//Ease the copy-pasta, this textarea is used to programmatically set clipboard
//content, see https://stackoverflow.com/a/30810322
var copyArea = document.getElementById("copyer");
function copyTextToClipboard(text) {
	copyArea.style.display = "block";
	copyArea.value = text;
	copyArea.select();
	try {
		var successful = document.execCommand('copy');
		if (successful){
			snack("Copied to clipboard", 3000);
		}else{
			snack("There was an error copying.", 3000);
		}
	} catch (err) {
		snack("There was an error copying " + err, 3000);
	}
	copyArea.style.display = "none";
}

var copyerjs = document.getElementById("copyerjs");
new Clipboard('#input-search', {
	text: function(trigger) {
		console.log("Triggered");
		return copyerjs;
		//copyerjs.display="none";
	}
});

//Handle click to copy
function listClick(event){
	var elm = event.target.parentElement.getElementsByClassName(ita?"ita":"eng");
	copyTextToClipboard(elm[0].textContent);
	searchbar.value = "";
	emptySearch();
	searchbar.focus();
	return false;
}

//Langsiwtch
var langbtn = document.getElementById("langbtn");
var ita = false;
function toggleLang(){
	ita = !ita;
	langbtn.innerText = ita?"ITA":"ENG";
}

//Snackbar
var snackbar = document.getElementById("snackbar");
function snack(text,time){
	snackbar.className = "show";
	snackbar.innerText = text;
	if (time>0){
		setTimeout( function(){
			snackbar.className = snackbar.className.replace("show", "");
		}, time);
	}
}

//Search logic
var searchbar = document.getElementById("input-search");
searchbar.focus()
var ws = new WebSocket("ws://localhost:42137/ws");
var reslist = document.getElementById("result-list");
function updateSearch(event){
	var msg = { "search" : searchbar.value };
	ws.send(JSON.stringify(msg));
};
ws.onopen = function (event) {
	console.log("Connected to backend")
	searchbar.addEventListener("input", updateSearch);
	console.log("Enabled search")
};
ws.onmessage = function (event) {
	var msg = JSON.parse(event.data);
	reslist.innerHTML = msg.Html;
};
ws.onclose = function(event){
	console.log("Server disconnected");
	searchbar.disabled = true;
	snack("Server closed the connection, please \nrealunch the application and reload the page", -1);
}
function emptySearch(){ reslist.innerHTML = ""; }

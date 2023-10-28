let menuicn = document.querySelector(".menuicn");
let nav = document.querySelector(".navcontainer");

menuicn.addEventListener("click", () => {
	nav.classList.toggle("navclose");
})

// Get the element with id="defaultOpen" and click on it
document.getElementById('defaultOpen').click();
// Get the element with id="defaultOps" and click on it
document.getElementById('defaultOps').click();

function openPage(pageName, elmnt, color) {
	// Hide all elements with class="tabcontent" by default */
	var i, tabcontent, tablink;
	tabcontent = document.getElementsByClassName("tabcontent");
	for (i = 0; i < tabcontent.length; i++) {
	  tabcontent[i].style.display = "none";
	}
  
	// Remove the background color of all tablinks/buttons
	tablink = document.getElementsByClassName("tablink");
	for (i = 0; i < tablink.length; i++) {
	  tablink[i].style.backgroundColor = "";
	}
  
	// Show the specific tab content
	document.getElementById(pageName).style.display = "flex";
  
	// Add the specific color to the button used to open the tab content
	elmnt.style.backgroundColor = color;
  }

function openFocus(focusName, elmnt, color) {
	// Hide all elements with class="tabcontent" by default */
	var i, tablink2, subtab;
	subtab = document.getElementsByClassName("subtab");
	for (i = 0; i < subtab.length; i++) {
		subtab[i].style.display = "none";
	}
	
	// Remove the background color of all tablinks/buttons
	tablink2 = document.getElementsByClassName("tablink2");
	for (i = 0; i < tablink2.length; i++) {
	  tablink2[i].style.backgroundColor = "";
	}
  
	// Show the specific tab content
	document.getElementById(focusName).style.display = "block";
  
	// Add the specific color to the button used to open the tab content
	elmnt.style.backgroundColor = color;
  }

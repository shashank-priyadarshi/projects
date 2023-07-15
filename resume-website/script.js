createCollapsibles();
function createCollapsibles() {
    let coll = document.getElementsByClassName("collapsible");
    let iter;
    for (iter = 0; iter < coll.length; iter++) {
        coll[iter].addEventListener("click", function () {
            if (!this.classList.contains("contact-disable")) {
                if (!this.classList.contains("sub-collapsible")) {
                    collapseOpen();
                } else {
                    let openSubCollapse = document.querySelectorAll(".sub-collapsible.active");
                    let subCollapseIter;
                    for (subCollapseIter = 0; subCollapseIter < openSubCollapse.length; subCollapseIter++) {
                        openSubCollapse[subCollapseIter].classList.remove("active");
                        openSubCollapse[subCollapseIter].nextElementSibling.style.display = "none";
                    }
                }
                this.classList.toggle("active");
                let content = this.nextElementSibling;
                if (content.style.display === "block") {
                    content.style.display = "none";
                } else {
                    content.style.display = "block";
                }
            } else {
                if (confirm("You are not authorised to access this content. Please drop a message on LinkedIn. You will now be redirected.")) {
                    window.location.href = document.getElementById("linkedin").href
                }
            }
            // if (content.style.maxHeight) {
            //     content.style.maxHeight = null;
            // } else {
            //     content.style.maxHeight = content.scrollHeight + "px";
            // }
        });
    }
}

function collapseOpen() {
    let open = document.querySelectorAll(".collapsible.active");
    let newiter;
    for (newiter = 0; newiter < open.length; newiter++) {
        open[newiter].classList.remove("active");
        open[newiter].nextElementSibling.style.display = "none";
    }
}

function openForm() {
    document.getElementById("github").style.display = "block";
    confirm("This part of the page is still under development.")
}

function closeForm() {
    document.getElementById("github").style.display = "none";
}
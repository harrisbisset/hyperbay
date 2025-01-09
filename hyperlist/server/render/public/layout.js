let current_event = "mouseout"

document.getElementById("prefix-hyper-header").addEventListener("mouseover", (e) => {
    if (current_event == "mouseout") {
        current_event = "mouseover"
        document.getElementById("service-postfix-hyper-header").style.opacity = "0"
        document.getElementById("bay-postfix-hyper-header").style.display = "block"
        document.getElementById("bay-postfix-hyper-header").style.opacity = "100"
        document.getElementById("hyper-header").style.cursor = "pointer"

    }
})

document.getElementById("hyper-header").addEventListener("mouseleave", (e) => {
    if (current_event == "mouseover") {
        current_event = "mouseout"
        document.getElementById("service-postfix-hyper-header").style.opacity = "100"
        document.getElementById("bay-postfix-hyper-header").style.opacity = "0"
        document.getElementById("bay-postfix-hyper-header").style.display = "none"
        document.getElementById("hyper-header").style.cursor = "default"
    }
})
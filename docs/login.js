const serverForm = document.querySelector("#serverForm")
const loginForm = document.querySelector("#loginForm")

const serverLog = document.querySelector("#serverLog")
const serverLogContent = document.querySelector("#serverLog label")

const httpPattern = /^http:\/\//

async function serverAssignment() {
    const serverData = new FormData(serverForm)
    server = serverData.get("server")

    if (!httpPattern.test(server)) {
        // Find :// then change to the beginning
    }

    serverLog.style.display = "none"
    serverLogContent.innerHTML = ""

    try {
        fetch(server)
        .then(response => {
            if (response.ok) {
                serverLog.style.display = "flex"
                serverLogContent.innerHTML = "Reachable server!"
            } else {
                serverLog.style.display = "flex"
                serverLogContent.innerHTML = "Server responded with an error"
            }
        })
    } catch (e) {
        console.error(e)
        serverLog.style.display = "flex"
        serverLogContent.innerHTML = e
    }
}

async function sendCredentials() {
    loginData = new FormData(loginForm)

    try {
        const response = await fetch(server + "/register", {
            method: "POST",
            body: JSON.stringify({
                "username": loginData.get("username"),
                "password": loginData.get("password")
            }),
            headers: {
                "Content-type": "application/json"
            }
        })
        registerResponse = await response.json()
        console.log(registerResponse)
    } catch (e) {
        console.error(e)
    }
    
    sendLogin()
}

async function sendLogin() {
    try {
        const response = await fetch(server + "/login", {
            method: "POST",
            body: JSON.stringify({
                "username": loginData.get("username"),
                "password": loginData.get("password")
            }),
            headers: {
                "Content-type": "application/json"
            }
        })
        loginResponse = await response.json()
        console.log(loginResponse)
    } catch (e) {
        console.error(e)
    }
}

function generateCookie(cname, cvalue, expdays) {
    const d = new Date()
    d.setTime(d.getTime() + (expdays*24*60*60*1000))
    let expires = "expires="+ d.toUTCString()
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/"
}

function login() {

}

serverForm.addEventListener("submit", (event) => {
    event.preventDefault()
    serverAssignment()
})

loginForm.addEventListener("submit", (event) => {
    event.preventDefault()
    //console.log(event.submitter)
    sendCredentials()
})

if (document.cookie) {
    login()
}
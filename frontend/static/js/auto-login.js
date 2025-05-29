async function validateTokenAndSetLink() {
    try {
        const res = await fetch("http://localhost:8081/validate", {
            method: 'GET',
            credentials: 'include',
        });

        const marketLink = document.getElementById("marketLink");

        if (res.ok) {
            marketLink.href = "/home";
            marketLink.classList.remove("disabled");
        }
    } catch (err) {
        console.log("Failed to connect to the server:", err);
    }
}

validateTokenAndSetLink();

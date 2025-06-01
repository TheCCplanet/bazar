const form = document.getElementById('login-form');
form.addEventListener('submit', async (e) => {
  e.preventDefault();

  const formData = new FormData(form);
  const payload = {
    username: formData.get('username'),
    password: formData.get('password'),
  };

  try {
    // auth-service/auth/login
    const res = await fetch('http://localhost:8081/auth', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include'
    });

    const data = await res.json();

    if (res.ok) {
      document.getElementById('result').innerText = data.message;
      window.location = data.redirect
    } else {
      document.getElementById('result').innerText = data.error || data.message;
    }
  } catch (err) {
    document.getElementById('result').innerText = 'Error connecting to auth service.';
  }
});
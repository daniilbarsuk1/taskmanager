document.addEventListener('DOMContentLoaded', () => 
    document.getElementById('form_login').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Array.from(formData.entries()).reduce((memo, [key, value]) => ({
            ...memo,
            [key]: value,
        }), {});
        let response = fetch(window.location.href, {
            redirect: 'follow',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            method: 'POST',
            body: JSON.stringify(data)
        }).then((response) => {
            if (response.redirected){
                alert("Регистрация завершена успешно!")
                location.href = response.url
            }
        })
    })
);

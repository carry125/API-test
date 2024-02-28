document.addEventListener("DOMContentLoaded", function () {
    const queryForm = document.getElementById("queryForm");
    const resultDiv = document.getElementById("result");

    queryForm.addEventListener("submit", function (e) {
        e.preventDefault();
        
        const formData = new FormData();
        formData.append("numStart", document.getElementById("start").value); //append("numStart" 此為API的PostForm指定字
        formData.append("numEnd", document.getElementById("end").value);

        // 发送API请求
        fetch("http://localhost:8080/isPrimesToApi", {
            method: "POST",
            body: formData
        })
        .then(response => response.json())
        .then(data => {  //data傳json出來 所以data底下的資料都會是json模樣
                if (data.error) {
                    resultDiv.innerHTML = `<p>Error: ${data.error}</p>`;
                } else {
                    // 显示素数表数据
                    resultDiv.innerHTML = `<h2>素数表：</h2>`;
                    const primeNumbers = data.prime_values; //golang "prime_values"
                    primeNumbers.forEach(number => {
                        resultDiv.innerHTML += `<p>${number}</p>`;
                    });
                }
            })
            .catch(error => {
                resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
            });
    });
});

const app = document.querySelector("#app");

window.addEventListener('load', () => {
    ListaUsuarios();
    
  });

const ListaUsuarios = async () => {
    return await fetch('/users', {
      method: 'GET',
    })
      .then(function (response) {
        //Leer Promesa
        return response.json();
      })
      .then(function (data) {
        /// Extraer datos a JSON
       
        data.forEach(function (user) {
          console.log(user);
        });
      });
  };
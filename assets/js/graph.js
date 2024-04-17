// Fonction pour mettre à jour les listes de vitesses et distances
var speed = [];
var distance = [];
var last6ofspeed = [];
var lastofdistance = [];
var lenData = 0;

// Création du deuxième graphique (line chart)
const ctx2 = document.getElementById('myChart2');
const chart2 = new Chart(ctx2, {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'km/h',
            data: [],
            borderWidth: 1
        }]
    },
    options: {
        scales: {
            y: {
                beginAtZero: true
            }
        }
    }
});

function getData(data) {
    if (data && Array.isArray(data)) {
        var newLenData = data.length;
        if (newLenData > lenData) {
            lenData = newLenData;
            var FloatSpeed = parseFloat(data[newLenData - 1].vitesse);
            var FloatDistance = parseFloat(data[newLenData - 1].distance);
            speed.push(FloatSpeed);
            distance.push(FloatDistance);
            if (speed.length > 6) {
                last6ofspeed = speed.slice(-6);
            } else {
                last6ofspeed = speed;
            }
            lastofdistance = distance.slice(-1);
            console.log(last6ofspeed, typeof(last6ofspeed));
            console.log(lastofdistance, typeof(lastofdistance));
            for (i = 0 ; i < speed.length ; i++) {
                console.log(speed[i]);
            }
            for (i = 0 ; i < distance.length ; i++) {
                console.log(distance[i]);
            }
            updateChart2(last6ofspeed);
        }
    } else {
        console.log('Les données du serveur sont null ou undefined.');
    }
}

function fetchDataAndUpdateLists() {
    // Effectuez une requête fetch pour récupérer les données du serveur
    fetch('/getData')
        .then(response => response.json())
        .then(data => {
            // Appelez la fonction pour mettre à jour les listes avec les données du robot
            getData(data);
        })
        .catch(error => console.error('Erreur lors de la récupération des données:', error));
}

// Appelez la fonction au chargement de la page pour afficher les données initiales
fetchDataAndUpdateLists();

// Rafraîchissement périodique toutes les 5 secondes (ajustez la fréquence selon vos besoins)
setInterval(fetchDataAndUpdateLists, 500);

// Fonction pour mettre à jour le deuxième graphique avec les nouvelles données de vitesses
function updateChart2(newData) {
    // Mettre à jour les données du deuxième graphique
    chart2.data.labels = Array.from({ length: newData.length }, (_, index) => index + 1);
    chart2.data.datasets[0].data = newData;

    // Appeler la méthode update pour mettre à jour le deuxième graphique
    chart2.update();
}

// Création du premier graphique (bar chart)
const ctx1 = document.getElementById('myChart1');
const chart1 = new Chart(ctx1, {
    type: 'bar',
    data: {
        labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
        datasets: [{
            label: '# of Votes',
            data: [12, 19, 3, 5, 2, 3],
            borderWidth: 1
        }]
    },
    options: {
        scales: {
            y: {
                beginAtZero: true
            }
        }
    }
});

// Création du troisième graphique (pie chart)
const ctx3 = document.getElementById('myChart3');
const chart3 = new Chart(ctx3, {
    type: 'pie',
    data: {
        labels: ['IA', 'Humain'],
        datasets: [{
            label: '# of Votes',
            data: [12, 19],
            borderWidth: 1
        }]
    },
});

import requests
import json
import time
import random

# Exemple de données à envoyer depuis le robot
robot_data = {
    'id_race': 1,
    'vitesse': 22.5,
    'distance': 10.2,
    'trix': 0.5,
    'triy': 0.5,
    'triz': 0.5,
}

# URL du serveur Go
url = 'http://127.0.0.1:3000/postData'  # Assurez-vous de remplacer par l'adresse correcte

# Envoi de la requête POST avec les données JSON
try:
    while True:
        randomidrace = random.randint(1, 5)
        randomvitesse = random.uniform(0, 100)
        randomdistance = random.uniform(0, 100)
        randomtrix = random.uniform(0, 30)
        randomtriy = random.uniform(0, 30)
        randomtriz = random.uniform(0, 30)
        robot_data = {
            'idrace': randomidrace, 
            'vitesse': randomvitesse,
            'distance': randomdistance,
            'tri_x': randomtrix,
            'tri_y': randomtriy,
            'tri_z': randomtriz,
        }
        response = requests.post(url, json=robot_data)
        response.raise_for_status()  # Lève une exception si la requête n'a pas abouti
        print('Données envoyées avec succès:', response.text)
        #attendre 1 seconde
        time.sleep(1)
except requests.exceptions.RequestException as err:
    print('Erreur lors de l\'envoi des données:', err)

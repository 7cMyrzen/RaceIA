function openModal(id) {
    document.getElementById(id).style.display = 'block';
    document.getElementById('overlay').style.display = 'block';
}

function closeModals() {
    var modals = document.getElementsByClassName('gmodal');
    for (var i = 0; i < modals.length; i++) {
        modals[i].style.display = 'none';
    }
    document.getElementById('overlay').style.display = 'none';
}

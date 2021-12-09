document.addEventListener("DOMContentLoaded", function(){
    const modal = document.getElementById('model-rule-explain');
    document.getElementById('modal-open').addEventListener('click', modalOpen);
    function modalOpen() {
    　　modal.style.display = 'block';
    };

    document.addEventListener('click', outsideClose);
    function outsideClose(e) {
    　　if (e.target == modal) {
    　　modal.style.display = 'none';
    　　};
    };

}, false);

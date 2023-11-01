document.addEventListener('DOMContentLoaded', function() {
  document.getElementById('toggleButton').addEventListener('click', function() {
      var targetDiv = document.getElementById('targetDiv');
      targetDiv.classList.toggle('hidden');

      var arrowIcon = document.getElementById('arrowIcon');
      arrowIcon.classList.toggle('fa-arrow-up');
      arrowIcon.classList.toggle('fa-arrow-down');

      var videoElements = document.getElementsByClassName('myVideos');

      for (var i = 0; i < videoElements.length; i++) {
          videoElements[i].play();
      }
  });
});


function showImage(selection) {
  var imageContainers = document.getElementsByClassName('image-container');
  for (var i = 0; i < imageContainers.length; i++) {
      imageContainers[i].style.display = 'none';
  }
  document.getElementById(selection + '-image').style.display = 'block';

  
  var selectors = document.getElementsByClassName('selectors')[0].children;
  for (var i = 0; i < selectors.length; i++) {
    selectors[i].classList.remove('selected');
  }

  
  document.querySelector('.' + selection).classList.add('selected');
}
  window.addEventListener('load', function() {
    showImage('card');
});


function toggleBioSections() {
  var bioSections = document.getElementsByClassName('bio-section');
  for (var i = 0; i < bioSections.length; i++) {
    bioSections[i].style.display = 'table-row';
  }

  var vaSections = document.getElementsByClassName('voiceActors-section');
  for (var i = 0; i < vaSections.length; i++) {
    vaSections[i].style.display = 'none';
  }

  document.querySelector('.Bio').classList.add('selected');
  document.querySelector('.VoiceActors').classList.remove('selected');
}

function toggleVASections() {
  var vaSections = document.getElementsByClassName('voiceActors-section');
  for (var i = 0; i < vaSections.length; i++) {
    vaSections[i].style.display = 'table-row';
  }

  var bioSections = document.getElementsByClassName('bio-section');
  for (var i = 0; i < bioSections.length; i++) {
    bioSections[i].style.display = 'none';
  }

  document.querySelector('.VoiceActors').classList.add('selected');
  document.querySelector('.Bio').classList.remove('selected');
}

document.addEventListener('DOMContentLoaded', function() {
  toggleBioSections();
});


document.addEventListener('DOMContentLoaded', function() {
  var scrollToTopButton = document.getElementById('scrollToTopButton');

  window.addEventListener('scroll', function() {
      if (window.scrollY > 300) { // Adjust this value as needed
          scrollToTopButton.style.display = 'block';
      } else {
          scrollToTopButton.style.display = 'none';
      }
  });

  scrollToTopButton.addEventListener('click', function() {
      window.scrollTo({
          top: 0,
          behavior: 'smooth' // This provides smooth scrolling
      });
  });
});
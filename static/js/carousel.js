const galleryContainer = document.querySelector('.gallery-container');
const galleryControlsContainer = document.querySelector('.gallery-controls');
const galleryControls = ['previous', 'next'];
const galleryItems = document.querySelectorAll('.gallery-item');
const Previous1 = '.gallery-item-first';
const Previous2 = '.gallery-item-previous';
const Next1 = '.gallery-item-next';
const Next2 = '.gallery-item-last';
let i = 0;

class Carousel {
  constructor(container, items, controls) {
    this.carouselContainer = container;
    this.carouselControls = controls;
    this.carouselArray = [...items];
  }

  // Assign initial css classes for gallery and nav items
  setInitialState() {
    this.carouselArray[0].classList.add('gallery-item-first');
    this.carouselArray[1].classList.add('gallery-item-previous');
    this.carouselArray[2].classList.add('gallery-item-selected');
    this.carouselArray[3].classList.add('gallery-item-next');
    this.carouselArray[4].classList.add('gallery-item-last');
    if (window.innerWidth > 1000) {
      document.querySelectorAll('.gallery-item .flex-item5').forEach(item => {
        item.disabled = true;
      });
    }
    document.querySelector('.gallery-item-selected .flex-item5').disabled = false;
    document.getElementsByClassName('gallery-item-selected')[0].getElementsByTagName('button')[0].classList.add('showBtn');
  }



  // Update the order state of the carousel with css classes
  setCurrentState(target, selected, previous, next, first, last) {

    selected.forEach(el => {
      el.classList.remove('gallery-item-selected');

      if (target.className == 'previousBtn') {
        el.classList.add('gallery-item-next');
      } else {
        el.classList.add('gallery-item-previous');
      }
    });

    previous.forEach(el => {
      el.classList.remove('gallery-item-previous');

      if (target.className == 'previousBtn') {
        el.classList.add('gallery-item-selected');
      } else {
        el.classList.add('gallery-item-first');
      }
    });

    next.forEach(el => {
      el.classList.remove('gallery-item-next');

      if (target.className == 'previousBtn') {
        el.classList.add('gallery-item-last');
      } else {
        el.classList.add('gallery-item-selected');
      }
    });

    first.forEach(el => {
      el.classList.remove('gallery-item-first');

      if (target.className == 'previousBtn') {
        el.classList.add('gallery-item-previous');
      } else {
        el.classList.add('gallery-item-last');
      }
    });

    last.forEach(el => {
      el.classList.remove('gallery-item-last');

      if (target.className == 'previousBtn') {
        el.classList.add('gallery-item-first');
      } else {
        el.classList.add('gallery-item-next');
      }
    });
    document.querySelectorAll('.gallery-item .flex-item5').forEach(item => {
      item.disabled = true;
    });
    useButtons();
    useXButton();
    document.querySelector('.gallery-item-selected .flex-item5').disabled = false;
    document.getElementsByClassName('gallery-item-selected')[0].getElementsByTagName('button')[0].classList.add('showBtn');
    document.getElementsByClassName('gallery-item-previous')[0].getElementsByTagName('button')[0].classList.remove('showBtn');
    document.getElementsByClassName('gallery-item-next')[0].getElementsByTagName('button')[0].classList.remove('showBtn');

    document.getElementsByClassName('gallery-item-selected')[0].getElementsByTagName('button')[0].classList.add('showBtn');
    document.querySelector('.gallery-item-first .div2').style.display = "none";
  }

  // Construct the carousel navigation
  setNav() {
    galleryContainer.appendChild(document.createElement('ul')).className = 'gallery-nav';

    this.carouselArray.forEach(item => {
      const nav = galleryContainer.lastElementChild;
      nav.appendChild(document.createElement('li'));
    });
  }

  // Construct the carousel controls
  setControls() {
    this.carouselControls.forEach(control => {
      galleryControlsContainer.appendChild(document.createElement('button')).className = `gallery-controls-${control}`;
    });

    !!galleryControlsContainer.childNodes[0] ? galleryControlsContainer.childNodes[0].innerHTML = null : null;
    !!galleryControlsContainer.childNodes[1] ? galleryControlsContainer.childNodes[1].innerHTML = null : null;
  }

  // Add a click event listener to trigger setCurrentState method to rearrange carousel
  useControls() {
    const triggers = [document.getElementsByClassName('previousBtn')[0], document.getElementsByClassName('nextBtn')[0]];

    triggers.forEach(control => {
      control.addEventListener('click', () => {
        const target = control;
        const selectedItem = document.querySelectorAll('.gallery-item-selected');
        const previousSelectedItem = document.querySelectorAll('.gallery-item-previous');
        const nextSelectedItem = document.querySelectorAll('.gallery-item-next');
        const firstCarouselItem = document.querySelectorAll('.gallery-item-first');
        const lastCarouselItem = document.querySelectorAll('.gallery-item-last');

        this.setCurrentState(target, selectedItem, previousSelectedItem, nextSelectedItem, firstCarouselItem, lastCarouselItem);
      });
    });
  }
}

const exampleCarousel = new Carousel(galleryContainer, galleryItems, galleryControls);

exampleCarousel.setControls();
exampleCarousel.setNav();
exampleCarousel.setInitialState();
exampleCarousel.useControls();

const buttons = [document.getElementsByClassName('previousBtn')[0], document.getElementsByClassName('nextBtn')[0]];
const items = [
  document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item1')[0],
  document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item2')[0],
  document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item3')[0],
  document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item4')[0],
  document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item5')[0]
];
const fade = [document.getElementsByClassName('gallery-item-first')[0],
document.getElementsByClassName('gallery-item-last')[0],
document.getElementsByClassName('gallery-item-previous')[0],
document.getElementsByClassName('gallery-item-next')[0]];

document.querySelector('.gallery-item-first .div2').style.display = "none";

a1 = document.querySelector('#a1');
a2 = document.querySelector('#a2');
a3 = document.querySelector('#a3');

function useButtons() {
  const controlbuttons = document.querySelectorAll('.gallery-item-selected .flex-item5');
  triggers = controlbuttons;

  triggers.forEach(button => {
    button.addEventListener('click', () => {
      const items = [
        document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item1')[0],
        document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item2')[0],
        document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item3')[0],
        document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item4')[0],
        document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item5')[0]
      ];
      const fade = [document.getElementsByClassName('gallery-item-first')[0],
      document.getElementsByClassName('gallery-item-last')[0],
      document.getElementsByClassName('gallery-item-previous')[0],
      document.getElementsByClassName('gallery-item-next')[0]];
      button.classList.add('opacity0');

      items[1].style.opacity = 0;
      items[2].style.opacity = 0;
      items[3].style.opacity = 0;
      document.querySelector('.gallery-item-selected .xButton').style.display = "inherit";
      document.querySelector('.gallery-item-selected .div2').style.display = "inherit";

      fade.forEach(fade => {
        fade.classList.add('opacity0');
        setTimeout(function () {
          fade.style.display = "none";
        }, 300);
      });
      buttons.forEach(item => {
        item.classList.add('opacity0');
      });

      setTimeout(function () {

        buttons.forEach(item => {
          item.style.display = "none";
        });
        button.style.display = "none";
        items[0].classList.add('img-change');
        document.querySelector('.gallery-item-selected .xButton').classList.remove('opacity0');

        items[1].classList.remove('transition');
        items[2].classList.remove('transition');
        items[3].classList.remove('transition');
        items[4].classList.remove('transition');
        items[1].style.display = "none";
        items[2].style.display = "none";
        items[3].style.display = "none";
      }, 300);

      document.querySelector('.gallery-item-selected .Line').style.display = 'inherit';
      document.querySelector('.gallery-item-selected .div1').style.display = 'inherit';
      document.querySelector('.gallery-item-selected .div2').style.display = 'inherit';
      document.querySelector('.gallery-item-selected .div3').style.display = 'inherit';
      document.querySelector('.gallery-item-selected .div4').style.display = 'inherit';
      document.querySelector('.gallery-item-selected .voteButton').style.display = 'inherit';
      button.parentElement.style.marginTop = "-50px";
      setTimeout(function () {
        document.querySelector('.gallery-item-selected .Line').classList.add('LineLong');
      }, 500);

      setTimeout(function () {
        button.parentElement.classList.add('boxBigger');
        document.querySelector('.gallery-item-selected .div1').style.opacity = 1;
        document.querySelector('.gallery-item-selected .div2').style.opacity = 1;
        document.querySelector('.gallery-item-selected .Line').style.opacity = 1;
        document.querySelector('.gallery-item-selected .div3').style.opacity = 1;
        document.querySelector('.gallery-item-selected .div4').style.opacity = 1;
        document.querySelector('.gallery-item-selected .voteButton').style.opacity = 1;
        document.querySelector('.gallery-item-selected .xButton').style.opacity = 1;
      }, 300);
    });
  });
}
function useButtons1() {
  const controlbuttons = document.querySelectorAll('.gallery-item-first .flex-item5');
  triggers = controlbuttons;

  triggers.forEach(button => {
    button.addEventListener('click', () => {
      const items = [
        document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item1')[0],
        document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item2')[0],
        document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item3')[0],
        document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item4')[0],
        document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item5')[0]
      ];
      const fade = [document.getElementsByClassName('gallery-item-selected')[0],
      document.getElementsByClassName('gallery-item-last')[0],
      document.getElementsByClassName('gallery-item-previous')[0],
      document.getElementsByClassName('gallery-item-next')[0]];
      button.classList.add('opacity0');

      items[1].style.opacity = 0;
      items[2].style.opacity = 0;
      items[3].style.opacity = 0;
      document.querySelector('.gallery-item-first .xButton').style.display = "inherit";
      document.querySelector('.gallery-item-first .div2').style.display = "inherit";

      fade.forEach(fade => {
        fade.classList.add('opacity0');
        setTimeout(function () {
          fade.style.display = "none";
        }, 300);
      });
      buttons.forEach(item => {
        item.classList.add('opacity0');
      });

      setTimeout(function () {

        buttons.forEach(item => {
          item.style.display = "none";
        });
        button.style.display = "none";
        items[0].classList.add('img-change');
        document.querySelector('.gallery-item-first .xButton').classList.remove('opacity0');

        items[1].classList.remove('transition');
        items[2].classList.remove('transition');
        items[3].classList.remove('transition');
        items[4].classList.remove('transition');
        items[1].style.display = "none";
        items[2].style.display = "none";
        items[3].style.display = "none";
      }, 300);

      document.querySelector('.gallery-item-first .Line').style.display = 'inherit';
      document.querySelector('.gallery-item-first .div1').style.display = 'inherit';
      document.querySelector('.gallery-item-first .div2').style.display = 'inherit';
      document.querySelector('.gallery-item-first .div3').style.display = 'inherit';
      document.querySelector('.gallery-item-first .div4').style.display = 'inherit';
      document.querySelector('.gallery-item-first .voteButton').style.display = 'inherit';
      button.parentElement.style.marginTop = "-50px";
      setTimeout(function () {
        document.querySelector('.gallery-item-first .Line').classList.add('LineLong');
      }, 500);

      setTimeout(function () {
        button.parentElement.classList.add('boxBigger');
        document.querySelector('.gallery-item-first .div1').style.opacity = 1;
        document.querySelector('.gallery-item-first .div2').style.opacity = 1;
        document.querySelector('.gallery-item-first .Line').style.opacity = 1;
        document.querySelector('.gallery-item-first .div3').style.opacity = 1;
        document.querySelector('.gallery-item-first .div4').style.opacity = 1;
        document.querySelector('.gallery-item-first .voteButton').style.opacity = 1;
        document.querySelector('.gallery-item-first .xButton').style.opacity = 1;
      }, 300);
    });
  });
}

function useButtons2() {
  const controlbuttons = document.querySelectorAll('.gallery-item-last .flex-item5');
  triggers = controlbuttons;

  triggers.forEach(button => {
    button.addEventListener('click', () => {
      const items = [
        document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item1')[0],
        document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item2')[0],
        document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item3')[0],
        document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item4')[0],
        document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item5')[0]
      ];
      const fade = [document.getElementsByClassName('gallery-item-selected')[0],
      document.getElementsByClassName('gallery-item-first')[0],
      document.getElementsByClassName('gallery-item-previous')[0],
      document.getElementsByClassName('gallery-item-next')[0]];
      button.classList.add('opacity0');

      items[1].style.opacity = 0;
      items[2].style.opacity = 0;
      items[3].style.opacity = 0;
      document.querySelector('.gallery-item-last .xButton').style.display = "inherit";
      document.querySelector('.gallery-item-last .div2').style.display = "inherit";

      fade.forEach(fade => {
        fade.classList.add('opacity0');
        setTimeout(function () {
          fade.style.display = "none";
        }, 300);
      });
      buttons.forEach(item => {
        item.classList.add('opacity0');
      });

      setTimeout(function () {

        buttons.forEach(item => {
          item.style.display = "none";
        });
        button.style.display = "none";
        items[0].classList.add('img-change');
        document.querySelector('.gallery-item-last .xButton').classList.remove('opacity0');

        items[1].classList.remove('transition');
        items[2].classList.remove('transition');
        items[3].classList.remove('transition');
        items[4].classList.remove('transition');
        items[1].style.display = "none";
        items[2].style.display = "none";
        items[3].style.display = "none";
      }, 300);

      document.querySelector('.gallery-item-last .Line').style.display = 'inherit';
      document.querySelector('.gallery-item-last .div1').style.display = 'inherit';
      document.querySelector('.gallery-item-last .div2').style.display = 'inherit';
      document.querySelector('.gallery-item-last .div3').style.display = 'inherit';
      document.querySelector('.gallery-item-last .div4').style.display = 'inherit';
      document.querySelector('.gallery-item-last .voteButton').style.display = 'inherit';
      button.parentElement.style.marginTop = "-50px";
      setTimeout(function () {
        document.querySelector('.gallery-item-last .Line').classList.add('LineLong');
      }, 500);

      setTimeout(function () {
        button.parentElement.classList.add('boxBigger');
        document.querySelector('.gallery-item-last .div1').style.opacity = 1;
        document.querySelector('.gallery-item-last .div2').style.opacity = 1;
        document.querySelector('.gallery-item-last .Line').style.opacity = 1;
        document.querySelector('.gallery-item-last .div3').style.opacity = 1;
        document.querySelector('.gallery-item-last .div4').style.opacity = 1;
        document.querySelector('.gallery-item-last .voteButton').style.opacity = 1;
        document.querySelector('.gallery-item-last .xButton').style.opacity = 1;
      }, 300);
    });
  });
}

function useButtons3() {
  const controlbuttons = document.querySelectorAll('.gallery-item-previous .flex-item5');
  triggers = controlbuttons;

  triggers.forEach(button => {
    button.addEventListener('click', () => {
      const items = [
        document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item1')[0],
        document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item2')[0],
        document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item3')[0],
        document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item4')[0],
        document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item5')[0]
      ];
      const fade = [document.getElementsByClassName('gallery-item-selected')[0],
      document.getElementsByClassName('gallery-item-last')[0],
      document.getElementsByClassName('gallery-item-first')[0],
      document.getElementsByClassName('gallery-item-next')[0]];
      button.classList.add('opacity0');

      items[1].style.opacity = 0;
      items[2].style.opacity = 0;
      items[3].style.opacity = 0;
      document.querySelector('.gallery-item-previous .xButton').style.display = "inherit";
      document.querySelector('.gallery-item-previous .div2').style.display = "inherit";

      fade.forEach(fade => {
        fade.classList.add('opacity0');
        setTimeout(function () {
          fade.style.display = "none";
        }, 300);
      });
      buttons.forEach(item => {
        item.classList.add('opacity0');
      });

      setTimeout(function () {

        buttons.forEach(item => {
          item.style.display = "none";
        });
        button.style.display = "none";
        items[0].classList.add('img-change');
        document.querySelector('.gallery-item-previous .xButton').classList.remove('opacity0');

        items[1].classList.remove('transition');
        items[2].classList.remove('transition');
        items[3].classList.remove('transition');
        items[4].classList.remove('transition');
        items[1].style.display = "none";
        items[2].style.display = "none";
        items[3].style.display = "none";
      }, 300);

      document.querySelector('.gallery-item-previous .Line').style.display = 'inherit';
      document.querySelector('.gallery-item-previous .div1').style.display = 'inherit';
      document.querySelector('.gallery-item-previous .div2').style.display = 'inherit';
      document.querySelector('.gallery-item-previous .div3').style.display = 'inherit';
      document.querySelector('.gallery-item-previous .div4').style.display = 'inherit';
      document.querySelector('.gallery-item-previous .voteButton').style.display = 'inherit';
      button.parentElement.style.marginTop = "-50px";
      setTimeout(function () {
        document.querySelector('.gallery-item-previous .Line').classList.add('LineLong');
      }, 500);

      setTimeout(function () {
        button.parentElement.classList.add('boxBigger');
        document.querySelector('.gallery-item-previous .div1').style.opacity = 1;
        document.querySelector('.gallery-item-previous .div2').style.opacity = 1;
        document.querySelector('.gallery-item-previous .Line').style.opacity = 1;
        document.querySelector('.gallery-item-previous .div3').style.opacity = 1;
        document.querySelector('.gallery-item-previous .div4').style.opacity = 1;
        document.querySelector('.gallery-item-previous .voteButton').style.opacity = 1;
        document.querySelector('.gallery-item-previous .xButton').style.opacity = 1;
      }, 300);
    });
  });
}

function useButtons4() {
  const controlbuttons = document.querySelectorAll('.gallery-item-next .flex-item5');
  triggers = controlbuttons;

  triggers.forEach(button => {
    button.addEventListener('click', () => {
      const items = [
        document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item1')[0],
        document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item2')[0],
        document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item3')[0],
        document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item4')[0],
        document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item5')[0]
      ];
      const fade = [document.getElementsByClassName('gallery-item-selected')[0],
      document.getElementsByClassName('gallery-item-last')[0],
      document.getElementsByClassName('gallery-item-previous')[0],
      document.getElementsByClassName('gallery-item-first')[0]];
      button.classList.add('opacity0');

      items[1].style.opacity = 0;
      items[2].style.opacity = 0;
      items[3].style.opacity = 0;
      document.querySelector('.gallery-item-next .xButton').style.display = "inherit";
      document.querySelector('.gallery-item-next .div2').style.display = "inherit";

      fade.forEach(fade => {
        fade.classList.add('opacity0');
        setTimeout(function () {
          fade.style.display = "none";
        }, 300);
      });
      buttons.forEach(item => {
        item.classList.add('opacity0');
      });

      setTimeout(function () {

        buttons.forEach(item => {
          item.style.display = "none";
        });
        button.style.display = "none";
        items[0].classList.add('img-change');
        document.querySelector('.gallery-item-next .xButton').classList.remove('opacity0');

        items[1].classList.remove('transition');
        items[2].classList.remove('transition');
        items[3].classList.remove('transition');
        items[4].classList.remove('transition');
        items[1].style.display = "none";
        items[2].style.display = "none";
        items[3].style.display = "none";
      }, 300);

      document.querySelector('.gallery-item-next .Line').style.display = 'inherit';
      document.querySelector('.gallery-item-next .div1').style.display = 'inherit';
      document.querySelector('.gallery-item-next .div2').style.display = 'inherit';
      document.querySelector('.gallery-item-next .div3').style.display = 'inherit';
      document.querySelector('.gallery-item-next .div4').style.display = 'inherit';
      document.querySelector('.gallery-item-next .voteButton').style.display = 'inherit';
      button.parentElement.style.marginTop = "-50px";
      setTimeout(function () {
        document.querySelector('.gallery-item-next .Line').classList.add('LineLong');
      }, 500);

      setTimeout(function () {
        button.parentElement.classList.add('boxBigger');
        document.querySelector('.gallery-item-next .div1').style.opacity = 1;
        document.querySelector('.gallery-item-next .div2').style.opacity = 1;
        document.querySelector('.gallery-item-next .Line').style.opacity = 1;
        document.querySelector('.gallery-item-next .div3').style.opacity = 1;
        document.querySelector('.gallery-item-next .div4').style.opacity = 1;
        document.querySelector('.gallery-item-next .voteButton').style.opacity = 1;
        document.querySelector('.gallery-item-next .xButton').style.opacity = 1;
      }, 300);
    });
  });
}


function useXButton() {
  const controlXbutton = document.querySelector('.gallery-item-selected .xButton');

  triggers = controlXbutton;


  triggers.addEventListener('click', () => {

    const items = [
      document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item1')[0],
      document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item2')[0],
      document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item3')[0],
      document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item4')[0],
      document.getElementsByClassName('gallery-item-selected')[0].getElementsByClassName('flex-item5')[0]
    ];
    const fade = [document.getElementsByClassName('gallery-item-first')[0],
    document.getElementsByClassName('gallery-item-last')[0],
    document.getElementsByClassName('gallery-item-previous')[0],
    document.getElementsByClassName('gallery-item-next')[0]];
    buttons.forEach(item => {
      item.style.display = "inherit";
    });
    fade.forEach(fade => {
      fade.style.display = "flex";
    });
    document.getElementsByClassName('gallery-item-selected')[0].classList.remove('boxBigger');
    document.querySelector('.gallery-item-selected .Line').classList.remove('LineLong');
    items[0].classList.remove('img-change');
    document.querySelector('.gallery-item-selected .div1').style.opacity = 0;
    document.querySelector('.gallery-item-selected .div2').style.opacity = 0;
    document.querySelector('.gallery-item-selected .Line').style.opacity = 0;
    document.querySelector('.gallery-item-selected .div3').style.opacity = 0;
    document.querySelector('.gallery-item-selected .div4').style.opacity = 0;
    document.querySelector('.gallery-item-selected .voteButton').style.opacity = 0;
    items[1].style.display = "block";
    items[2].style.display = "block";
    items[3].style.display = "block";
    items[4].style.display = "block";
    items[1].classList.add('transition');
    items[2].classList.add('transition');
    items[3].classList.add('transition');
    items[4].classList.add('transition');
    items[1].style.opacity = 0;
    items[2].style.opacity = 0;
    items[3].style.opacity = 0;

    document.querySelector('.gallery-item-selected .xButton').classList.add('opacity0');
    setTimeout(function () {
      items[1].style.opacity = 1;
      items[2].style.opacity = 1;
      items[3].style.opacity = 1;
      items[4].classList.remove('opacity0');
      document.querySelector('.gallery-item-selected .div1').style.display = "none";
      document.querySelector('.gallery-item-selected .div2').style.display = "none";
      document.querySelector('.gallery-item-selected .Line').style.display = "none";
      document.querySelector('.gallery-item-selected .div3').style.display = "none";
      document.querySelector('.gallery-item-selected .div4').style.display = "none";
      document.querySelector('.gallery-item-selected .voteButton').style.display = "none";
      document.querySelector('.gallery-item-selected .xButton').style.display = "none";
      document.getElementsByClassName('gallery-item-selected')[0].style.marginTop = 0;
      fade.forEach(fade => {
        fade.classList.remove('opacity0');
      });
      buttons.forEach(item => {
        item.classList.remove('opacity0');
      });
    }, 300);

    setTimeout(function () {
      items[1].classList.remove('transition');
      items[2].classList.remove('transition');
      items[3].classList.remove('transition');
      items[4].classList.remove('transition');
    }, 400);
  });
}

function useXButton1() {
  const controlXbutton = document.querySelector('.gallery-item-first .xButton');

  triggers = controlXbutton;


  triggers.addEventListener('click', () => {

    const items = [
      document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item1')[0],
      document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item2')[0],
      document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item3')[0],
      document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item4')[0],
      document.getElementsByClassName('gallery-item-first')[0].getElementsByClassName('flex-item5')[0]
    ];
    const fade = [document.getElementsByClassName('gallery-item-selected')[0],
    document.getElementsByClassName('gallery-item-last')[0],
    document.getElementsByClassName('gallery-item-previous')[0],
    document.getElementsByClassName('gallery-item-next')[0]];
    buttons.forEach(item => {
      item.style.display = "inherit";
    });
    fade.forEach(fade => {
      fade.style.display = "flex";
    });
    document.getElementsByClassName('gallery-item-first')[0].classList.remove('boxBigger');
    document.querySelector('.gallery-item-first .Line').classList.remove('LineLong');
    items[0].classList.remove('img-change');
    document.querySelector('.gallery-item-first .div1').style.opacity = 0;
    document.querySelector('.gallery-item-first .div2').style.opacity = 0;
    document.querySelector('.gallery-item-first .Line').style.opacity = 0;
    document.querySelector('.gallery-item-first .div3').style.opacity = 0;
    document.querySelector('.gallery-item-first .div4').style.opacity = 0;
    document.querySelector('.gallery-item-first .voteButton').style.opacity = 0;
    items[1].style.display = "block";
    items[2].style.display = "block";
    items[3].style.display = "block";
    items[4].style.display = "block";
    items[1].classList.add('transition');
    items[2].classList.add('transition');
    items[3].classList.add('transition');
    items[4].classList.add('transition');
    items[1].style.opacity = 0;
    items[2].style.opacity = 0;
    items[3].style.opacity = 0;

    document.querySelector('.gallery-item-first .xButton').classList.add('opacity0');
    setTimeout(function () {
      items[1].style.opacity = 1;
      items[2].style.opacity = 1;
      items[3].style.opacity = 1;
      items[4].classList.remove('opacity0');
      document.querySelector('.gallery-item-first .div1').style.display = "none";
      document.querySelector('.gallery-item-first .div2').style.display = "none";
      document.querySelector('.gallery-item-first .Line').style.display = "none";
      document.querySelector('.gallery-item-first .div3').style.display = "none";
      document.querySelector('.gallery-item-first .div4').style.display = "none";
      document.querySelector('.gallery-item-first .voteButton').style.display = "none";
      document.querySelector('.gallery-item-first .xButton').style.display = "none";
      document.getElementsByClassName('gallery-item-first')[0].style.marginTop = 0;
      fade.forEach(fade => {
        fade.classList.remove('opacity0');
      });
      buttons.forEach(item => {
        item.classList.remove('opacity0');
      });
    }, 300);

    setTimeout(function () {
      items[1].classList.remove('transition');
      items[2].classList.remove('transition');
      items[3].classList.remove('transition');
      items[4].classList.remove('transition');
    }, 400);
  });
}

function useXButton4() {
  const controlXbutton = document.querySelector('.gallery-item-next .xButton');

  triggers = controlXbutton;


  triggers.addEventListener('click', () => {

    const items = [
      document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item1')[0],
      document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item2')[0],
      document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item3')[0],
      document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item4')[0],
      document.getElementsByClassName('gallery-item-next')[0].getElementsByClassName('flex-item5')[0]
    ];
    const fade = [document.getElementsByClassName('gallery-item-selected')[0],
    document.getElementsByClassName('gallery-item-last')[0],
    document.getElementsByClassName('gallery-item-previous')[0],
    document.getElementsByClassName('gallery-item-first')[0]];
    buttons.forEach(item => {
      item.style.display = "inherit";
    });
    fade.forEach(fade => {
      fade.style.display = "flex";
    });
    document.getElementsByClassName('gallery-item-next')[0].classList.remove('boxBigger');
    document.querySelector('.gallery-item-next .Line').classList.remove('LineLong');
    items[0].classList.remove('img-change');
    document.querySelector('.gallery-item-next .div1').style.opacity = 0;
    document.querySelector('.gallery-item-next .div2').style.opacity = 0;
    document.querySelector('.gallery-item-next .Line').style.opacity = 0;
    document.querySelector('.gallery-item-next .div3').style.opacity = 0;
    document.querySelector('.gallery-item-next .div4').style.opacity = 0;
    document.querySelector('.gallery-item-next .voteButton').style.opacity = 0;
    items[1].style.display = "block";
    items[2].style.display = "block";
    items[3].style.display = "block";
    items[4].style.display = "block";
    items[1].classList.add('transition');
    items[2].classList.add('transition');
    items[3].classList.add('transition');
    items[4].classList.add('transition');
    items[1].style.opacity = 0;
    items[2].style.opacity = 0;
    items[3].style.opacity = 0;

    document.querySelector('.gallery-item-next .xButton').classList.add('opacity0');
    setTimeout(function () {
      items[1].style.opacity = 1;
      items[2].style.opacity = 1;
      items[3].style.opacity = 1;
      items[4].classList.remove('opacity0');
      document.querySelector('.gallery-item-next .div1').style.display = "none";
      document.querySelector('.gallery-item-next .div2').style.display = "none";
      document.querySelector('.gallery-item-next .Line').style.display = "none";
      document.querySelector('.gallery-item-next .div3').style.display = "none";
      document.querySelector('.gallery-item-next .div4').style.display = "none";
      document.querySelector('.gallery-item-next .voteButton').style.display = "none";
      document.querySelector('.gallery-item-next .xButton').style.display = "none";
      document.getElementsByClassName('gallery-item-next')[0].style.marginTop = 0;
      fade.forEach(fade => {
        fade.classList.remove('opacity0');
      });
      buttons.forEach(item => {
        item.classList.remove('opacity0');
      });
    }, 300);

    setTimeout(function () {
      items[1].classList.remove('transition');
      items[2].classList.remove('transition');
      items[3].classList.remove('transition');
      items[4].classList.remove('transition');
    }, 400);
  });
}

function useXButton2() {
  const controlXbutton = document.querySelector('.gallery-item-last .xButton');

  triggers = controlXbutton;


  triggers.addEventListener('click', () => {

    const items = [
      document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item1')[0],
      document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item2')[0],
      document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item3')[0],
      document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item4')[0],
      document.getElementsByClassName('gallery-item-last')[0].getElementsByClassName('flex-item5')[0]
    ];
    const fade = [document.getElementsByClassName('gallery-item-selected')[0],
    document.getElementsByClassName('gallery-item-first')[0],
    document.getElementsByClassName('gallery-item-previous')[0],
    document.getElementsByClassName('gallery-item-next')[0]];
    buttons.forEach(item => {
      item.style.display = "inherit";
    });
    fade.forEach(fade => {
      fade.style.display = "flex";
    });
    document.getElementsByClassName('gallery-item-last')[0].classList.remove('boxBigger');
    document.querySelector('.gallery-item-last .Line').classList.remove('LineLong');
    items[0].classList.remove('img-change');
    document.querySelector('.gallery-item-last .div1').style.opacity = 0;
    document.querySelector('.gallery-item-last .div2').style.opacity = 0;
    document.querySelector('.gallery-item-last .Line').style.opacity = 0;
    document.querySelector('.gallery-item-last .div3').style.opacity = 0;
    document.querySelector('.gallery-item-last .div4').style.opacity = 0;
    document.querySelector('.gallery-item-last .voteButton').style.opacity = 0;
    items[1].style.display = "block";
    items[2].style.display = "block";
    items[3].style.display = "block";
    items[4].style.display = "block";
    items[1].classList.add('transition');
    items[2].classList.add('transition');
    items[3].classList.add('transition');
    items[4].classList.add('transition');
    items[1].style.opacity = 0;
    items[2].style.opacity = 0;
    items[3].style.opacity = 0;

    document.querySelector('.gallery-item-last .xButton').classList.add('opacity0');
    setTimeout(function () {
      items[1].style.opacity = 1;
      items[2].style.opacity = 1;
      items[3].style.opacity = 1;
      items[4].classList.remove('opacity0');
      document.querySelector('.gallery-item-last .div1').style.display = "none";
      document.querySelector('.gallery-item-last .div2').style.display = "none";
      document.querySelector('.gallery-item-last .Line').style.display = "none";
      document.querySelector('.gallery-item-last .div3').style.display = "none";
      document.querySelector('.gallery-item-last .div4').style.display = "none";
      document.querySelector('.gallery-item-last .voteButton').style.display = "none";
      document.querySelector('.gallery-item-last .xButton').style.display = "none";
      document.getElementsByClassName('gallery-item-last')[0].style.marginTop = 0;
      fade.forEach(fade => {
        fade.classList.remove('opacity0');
      });
      buttons.forEach(item => {
        item.classList.remove('opacity0');
      });
    }, 300);

    setTimeout(function () {
      items[1].classList.remove('transition');
      items[2].classList.remove('transition');
      items[3].classList.remove('transition');
      items[4].classList.remove('transition');
    }, 400);
  });
}

function useXButton3() {
  const controlXbutton = document.querySelector('.gallery-item-previous .xButton');

  triggers = controlXbutton;


  triggers.addEventListener('click', () => {

    const items = [
      document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item1')[0],
      document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item2')[0],
      document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item3')[0],
      document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item4')[0],
      document.getElementsByClassName('gallery-item-previous')[0].getElementsByClassName('flex-item5')[0]
    ];
    const fade = [document.getElementsByClassName('gallery-item-selected')[0],
    document.getElementsByClassName('gallery-item-last')[0],
    document.getElementsByClassName('gallery-item-first')[0],
    document.getElementsByClassName('gallery-item-next')[0]];
    buttons.forEach(item => {
      item.style.display = "inherit";
    });
    fade.forEach(fade => {
      fade.style.display = "flex";
    });
    document.getElementsByClassName('gallery-item-previous')[0].classList.remove('boxBigger');
    document.querySelector('.gallery-item-previous .Line').classList.remove('LineLong');
    items[0].classList.remove('img-change');
    document.querySelector('.gallery-item-previous .div1').style.opacity = 0;
    document.querySelector('.gallery-item-previous .div2').style.opacity = 0;
    document.querySelector('.gallery-item-previous .Line').style.opacity = 0;
    document.querySelector('.gallery-item-previous .div3').style.opacity = 0;
    document.querySelector('.gallery-item-previous .div4').style.opacity = 0;
    document.querySelector('.gallery-item-previous .voteButton').style.opacity = 0;
    items[1].style.display = "block";
    items[2].style.display = "block";
    items[3].style.display = "block";
    items[4].style.display = "block";
    items[1].classList.add('transition');
    items[2].classList.add('transition');
    items[3].classList.add('transition');
    items[4].classList.add('transition');
    items[1].style.opacity = 0;
    items[2].style.opacity = 0;
    items[3].style.opacity = 0;

    document.querySelector('.gallery-item-previous .xButton').classList.add('opacity0');
    setTimeout(function () {
      items[1].style.opacity = 1;
      items[2].style.opacity = 1;
      items[3].style.opacity = 1;
      items[4].classList.remove('opacity0');
      document.querySelector('.gallery-item-previous .div1').style.display = "none";
      document.querySelector('.gallery-item-previous .div2').style.display = "none";
      document.querySelector('.gallery-item-previous .Line').style.display = "none";
      document.querySelector('.gallery-item-previous .div3').style.display = "none";
      document.querySelector('.gallery-item-previous .div4').style.display = "none";
      document.querySelector('.gallery-item-previous .voteButton').style.display = "none";
      document.querySelector('.gallery-item-previous .xButton').style.display = "none";
      document.getElementsByClassName('gallery-item-previous')[0].style.marginTop = 0;
      fade.forEach(fade => {
        fade.classList.remove('opacity0');
      });
      buttons.forEach(item => {
        item.classList.remove('opacity0');
      });
    }, 300);

    setTimeout(function () {
      items[1].classList.remove('transition');
      items[2].classList.remove('transition');
      items[3].classList.remove('transition');
      items[4].classList.remove('transition');
    }, 400);
  });
}

function voteCheck1(number) {
  if (confirm("한번 투표하면 번복은 불가합니다. 투표하시겠습니까?") == true) {    //확인
    location.href = "/api/voteCandidate?field=1&studentNumberCandidate=" + number;
  } else {   //취소
    return false;
  }
}

function voteCheck2(number) {
  if (confirm("한번 투표하면 번복은 불가합니다. 투표하시겠습니까?") == true) {    //확인
    location.href = "/api/voteCandidate?field=2&studentNumberCandidate=" + number;
  } else {   //취소
    return false;
  }
}
if (window.innerWidth > 1000) {
useButtons();
useXButton();
}
else{
  useButtons();
  useButtons1();
  useButtons2();
  useButtons3();
  useButtons4();
  useXButton1();
  useXButton2();
  useXButton3();
  useXButton4();
}



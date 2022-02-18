import './style.css';

const startup = () => {
  const video = document.querySelector('#videoElement');

  if (navigator.mediaDevices.getUserMedia) {
    navigator.mediaDevices
      .getUserMedia({
        video: true,
      })
      .then((stream) => {
        video.srcObject = stream;
      })
      .catch((err) => {
        console.error(err);
      });
  }
};

startup();

function isMobileDevice() {
    return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
}

const hubBody = document.getElementById('hubBody');
let mouseX = window.innerWidth / 20;
let mouseY = window.innerHeight / 20;
let targetX = mouseX;
let targetY = mouseY;
const sensitivity = 0.005;
const maxDeltaX = 30;
const maxDeltaY = 11;
const maxHubX = 300;
const minHubX = -300;
const maxHubY = 150;
const minHubY = -150;

let isMouseInside = false;
let mouseTimeout;

function resetHubPosition() {
    mouseX = window.innerWidth / 20;
    mouseY = window.innerHeight / 20;
    targetX = mouseX;
    targetY = mouseY;
    hubBody.style.transition = 'transform 0.5s ease-in-out';
    hubBody.style.transform = `translate(${mouseX}px, ${mouseY}px)`;
}

document.addEventListener("DOMContentLoaded", function () {
    var overlay = document.querySelector(".overlay");
    var loadbar = document.querySelector(".loadbar");
    var upper = document.querySelector(".upper");
    var lower = document.querySelector(".lower");

    setTimeout(function () {
        loadbar.style.width = "100%";
        loadbar.style.opacity = "1";
    }, 1);

    setTimeout(function () {
        upper.style.borderBottom = "1px solid #90f7e8";
        lower.style.borderTop = "1px solid #90f7e8";
        loadbar.style.opacity = "0";
    }, 1000);

    setTimeout(function () {
        overlay.style.opacity = "0";
    }, 1500);
});

function updateSpeedValue() {
    const speedElement = document.getElementById('speedValue');
    let currentSpeed = parseFloat(speedElement.innerText.replace(/,/g, ''));
    const targetSpeed = Math.floor(Math.random() * (1530000 - 1400000 + 1)) + 1400000;
    currentSpeed += (targetSpeed - currentSpeed) * 0.001;
    speedElement.innerText = currentSpeed.toLocaleString();
}

function updateAltValue() {
    const altElement = document.getElementById('altValue');
    let currentAlt = parseFloat(altElement.innerText.replace(/,/g, ''));
    const targetAlt = 10000;
    currentAlt -= (currentAlt - targetAlt) * 0.0001;
    altElement.innerText = currentAlt.toFixed(2);
}

function updateParallaxBackground() {
    const parallaxBackground = document.querySelector('.home-background');
    const parallaxSpeed = 0.4; // Ajustez la vitesse de parallaxe selon vos préférences

    if (window.innerWidth > 600) {
        const backgroundX = -mouseX * parallaxSpeed + window.innerWidth / 2;
        const backgroundY = -mouseY * parallaxSpeed + window.innerHeight / 2;

        const backgroundCenterX = backgroundX - parallaxBackground.clientWidth / 1.2;
        const backgroundCenterY = backgroundY - parallaxBackground.clientHeight / 1.2;

        parallaxBackground.style.backgroundPosition = `${backgroundCenterX}px ${backgroundCenterY}px`;
    } else {
        
    }
}


function updateMoonParallax() {
    const moonBackground = document.querySelector('.moon-background');
    const parallaxSpeed = 0.75; // Ajustez la vitesse de parallaxe selon vos préférences

    const backgroundX = -mouseX * parallaxSpeed + window.innerWidth / 2;
    const backgroundY = -mouseY * parallaxSpeed + window.innerHeight / 2;

    // Ajustez les coordonnées du centre du fond d'écran
    const backgroundCenterX = backgroundX - moonBackground.clientWidth / 4.8;
    const backgroundCenterY = backgroundY - moonBackground.clientHeight / 8;

    moonBackground.style.backgroundPosition = `${backgroundCenterX}px ${backgroundCenterY}px`;
}

function updateHubPosition() {
    const deltaX = targetX - window.innerWidth / 2;
    const deltaY = targetY - window.innerHeight / 2;

    mouseX += Math.min(maxDeltaX, Math.max(-maxDeltaX, deltaX)) * sensitivity;
    mouseY += Math.min(maxDeltaY, Math.max(-maxDeltaY, deltaY)) * sensitivity;

    mouseX = Math.min(maxHubX, Math.max(minHubX, mouseX));
    mouseY = Math.min(maxHubY, Math.max(minHubY, mouseY));

    hubBody.style.transform = `translate(${mouseX}px, ${mouseY}px)`;
    updateParallaxBackground();
    updateMoonParallax();
}

hubBody.addEventListener('mouseenter', function () {
    isMouseInside = true;
    clearTimeout(mouseTimeout);
    hubBody.style.transition = '';
});

hubBody.addEventListener('mouseleave', function () {
    isMouseInside = false;
    mouseTimeout = setTimeout(() => {
        if (!isMouseInside) {
            resetHubPosition();
        }
    }, 2000);
});

document.addEventListener('click', function (event) {
    if (event.button === 0) {
        createBullet(event.clientY);
    }
});

function createBullet(mouseY) {
    const bullet = document.createElement('div');
    bullet.className = 'bullet';
    document.body.appendChild(bullet);

    const hubRect = hubBody.getBoundingClientRect(); 

    const bulletSpeed = 90;
    const bulletPosition = { x: window.innerWidth, y: mouseY - bullet.offsetHeight / 2 };
    const bulletTarget = { x: window.innerWidth / 2, y: window.innerHeight / 2 };

    let canShoot = true;

    function updateBullet() {
        const deltaX = bulletTarget.x - bulletPosition.x;
        const deltaY = bulletTarget.y - bulletPosition.y;
        const distance = Math.sqrt(deltaX * deltaX + deltaY * deltaY);

        const bulletAngle = Math.atan2(deltaY, deltaX);
        const bulletXSpeed = bulletSpeed * Math.cos(bulletAngle);
        const bulletYSpeed = bulletSpeed * Math.sin(bulletAngle);

        if (distance > bulletSpeed) {
            bulletPosition.x += bulletXSpeed;
            bulletPosition.y += bulletYSpeed;
            bullet.style.left = bulletPosition.x + 'px';
            bullet.style.top = bulletPosition.y + 'px';

            const scale = 10 * distance / hubRect.width;
            bullet.style.transform = `scale(${Math.max(scale, 0)})`;

            requestAnimationFrame(updateBullet);
        } else {
            if (document.body.contains(bullet)) {
                document.body.removeChild(bullet);
            }
            canShoot = true;
        }
    }

    if (!canShoot) {
        return;
    }

    //removeBullets();
    canShoot = false;

    setTimeout(() => {
        updateBullet();
    }, 1);
}

function removeBullets() {
    const bullets = document.querySelectorAll('.bullet');

    bullets.forEach(bullet => {
        bullet.parentNode.removeChild(bullet);
    });
}

document.addEventListener('click', (event) => {
    createBullet(event.clientY);
});

let canShootLazer = true;

function createDualLazers() {
    createLazer('left');
    createLazer('right');
}

document.addEventListener('keydown', function (event) {
    if (event.key === 'r' && canShootLazer) {
        createDualLazers();
    }
});

function createLazer(side) {
    const lazer = document.createElement('div');
    lazer.className = 'lazer';
    document.body.appendChild(lazer);

    const lazerSpeed = 50;
    const lazerPosition = {
        x: side === 'left' ? 0 : window.innerWidth,
        y: window.innerHeight / 2 - lazer.offsetHeight / 2
    };

    const hubRect = hubBody.getBoundingClientRect();
    const lazerTarget = {
        x: hubRect.left + hubRect.width / 2,
        y: hubRect.top + hubRect.height / 2
    };

    function updateLazer() {
        const deltaX = lazerTarget.x - lazerPosition.x;
        const deltaY = lazerTarget.y - lazerPosition.y;
        const distance = Math.sqrt(deltaX * deltaX + deltaY * deltaY);

        const lazerAngle = Math.atan2(deltaY, deltaX);
        const lazerXSpeed = lazerSpeed * Math.cos(lazerAngle);
        const lazerYSpeed = lazerSpeed * Math.sin(lazerAngle);

        if (distance > lazerSpeed) {
            lazerPosition.x += lazerXSpeed;
            lazerPosition.y += lazerYSpeed;
            lazer.style.left = lazerPosition.x + 'px';
            lazer.style.top = lazerPosition.y + 'px';

            const scale = 10 * distance / window.innerWidth;
            lazer.style.transform = `scaleX(${Math.max(scale, 0)})`;

            requestAnimationFrame(updateLazer);
        } else {
            if (document.body.contains(lazer)) {
                document.body.removeChild(lazer);
            }
            canShootLazer = true;
        }
    }

    canShootLazer = false;

    setTimeout(() => {
        updateLazer();
    }, 1);
}

let railgunKey = document.querySelector('.railgun-key');
let railgunPicto = document.querySelector('.railgun-picto');

document.addEventListener('mousedown', function (event) {
    if (event.button === 0) { // Vérifiez si le bouton enfoncé est le bouton gauche (0)
        railgunKey.style.transition = 'background-color 0.4s, color 0.2s, border 0.4s, opacity 0.4s';
        railgunKey.style.backgroundColor = 'rgba(0, 0, 0, 0.9)';
        railgunKey.style.color = '#FF4500';
        railgunKey.style.border = '1px ridge #FF4500';
        railgunKey.style.opacity = '0.9';

        railgunPicto.style.transition = 'opacity 0.5s';
        railgunPicto.style.opacity = '0.9';
    }
});
document.addEventListener('mouseup', function (event) {
    if (event.button === 0) { // Vérifiez si le bouton relâché est le bouton gauche (0)
        railgunKey.style.transition = 'background-color 3s, color 3s, border 3s, opacity s';
        railgunKey.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
        railgunKey.style.color = '#90f7e8';
        railgunKey.style.border = '1px ridge #90f7e8';
        railgunKey.style.opacity = '0.65';
        railgunPicto.style.transition = 'opacity 2s';
        railgunPicto.style.opacity = '0.4';
    }
});


let rayKey = document.querySelector('.ray-key');
let rayPicto = document.querySelector('.ray-picto');

document.addEventListener('keydown', function (event) {
    if (event.key === 'r') {
        rayKey.style.transition = 'background-color 0.4s, color 0.2s, border 0.4s, opacity 0.4s';
        rayKey.style.backgroundColor = 'rgba(0, 0, 0, 0.9)';
        rayKey.style.color = '#FF4500';
        rayKey.style.border = '1px ridge #FF4500';
        rayKey.style.opacity = '0.9';

        rayPicto.style.transition = 'opacity 0.5s';
        rayPicto.style.opacity = '0.9';
    }
});
document.addEventListener('keyup', function (event) {
    if (event.key === 'r') {
        rayKey.style.transition = 'background-color 3s, color 3s, border 3s, opacity s';
        rayKey.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
        rayKey.style.color = '#90f7e8';
        rayKey.style.border = '1px ridge #90f7e8';
        rayKey.style.opacity = '0.65';

        rayPicto.style.transition = 'opacity 2s';
        rayPicto.style.opacity = '0.4';
    }
});

if (isMobileDevice()) {
    window.addEventListener('deviceorientation', function (event) {
        const deltaX = event.gamma; // Inclinaison à gauche/droite
        const deltaY = event.beta; // Inclinaison avant/arrière

        mouseX = window.innerWidth / 2 + deltaX * sensitivity * 50;
        mouseY = window.innerHeight / 2 + deltaY * sensitivity * 50;

        updateHubPosition();
    });
} else {
    document.addEventListener('mousemove', function (event) {
        targetX = event.clientX;
        targetY = event.clientY;
    });
}

setInterval(() => {
    updateSpeedValue();
}, 800);

setInterval(() => {
    updateAltValue();
}, 200);

setInterval(() => {
    updateHubPosition();
}, 16);


const alertMessage = document.getElementById('alertMessage');
        let isShieldActive = false;

        document.addEventListener('keydown', function(event) {
            let message = '';

            switch(event.key.toUpperCase()) {
                case 'M':
                    console.log('M key pressed');
                    message = 'We are out of missile';
                    break;
                case 'B':
                    message = 'We are out of bomb';
                    break;
                case 'S':
                    if (isShieldActive) {
                        message = 'Shield disabled';
                    } else {
                        message = 'Shield active';
                    }
                    isShieldActive = !isShieldActive;
                    break;
            }

            if (message !== '') {
                alertMessage.innerText = message;
                alertMessage.style.opacity = '1';

                setTimeout(() => {
                    alertMessage.style.opacity = '0';
                }, 5000); // Afficher pendant 5 secondes
            }
        });
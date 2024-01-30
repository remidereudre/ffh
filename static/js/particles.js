document.addEventListener('DOMContentLoaded', function () {
    const particles = [];

    function generateParticle() {
        const speed = 2; 
        const angle = Math.random() * Math.PI * 2; 

        particles.push({
            x: window.innerWidth / 1.6,
            y: window.innerHeight / 1.9,
            vx: Math.cos(angle) * speed,
            vy: Math.sin(angle) * speed,
            size: 2, 
            opacity: 0, 
        });
    }

    function drawParticles() {
        const canvas = document.querySelector('#particles-js canvas');
        const context = canvas.getContext('2d');

        context.clearRect(0, 0, window.innerWidth * 2, window.innerHeight * 2);

        particles.forEach((particle, index) => {
            context.beginPath();
            context.arc(particle.x, particle.y, particle.size, 0, Math.PI * 2);
            context.fillStyle = `rgba(255, 255, 255, ${particle.opacity})`;
            context.fill();

            particle.x += particle.vx;
            particle.y += particle.vy;

            // Augmenter l'opacité au fil du temps
            particle.opacity += 0.002;

            // Supprimer les particules lorsque leur opacité est suffisamment élevée
            if (particle.opacity >= 1.0) {
                particles.splice(index, 1);
                generateParticle();
            }
        });

        requestAnimationFrame(drawParticles);
    }

    // Générer initialement quelques particules
    for (let i = 0; i < 5; i++) {
        //generateParticle();
    }

    // Utiliser setInterval pour générer une particule toutes les 1000 millisecondes (1 seconde)
    const particleInterval = setInterval(generateParticle, 5000);

    // Arrêter la génération de particules après un certain temps (par exemple, 10 secondes)
    setTimeout(() => {
        clearInterval(particleInterval);
    }, 10000);

    particlesJS('particles-js', {
        particles: {
            number: { value: particles.length },
            color: { value: '#D2B48C' },
            shape: { type: 'circle' },
            opacity: { value: 0.01, random: true },
            size: { value: 2 }, // Taille fixe
            move: {
                enable: false,
            },
            line_linked: { enable: false },
        },
        interactivity: {
            detect_on: 'canvas',
            events: { onhover: { enable: true, mode: 'repulse' } },
        },
        retina_detect: true,
    });

    drawParticles();
});

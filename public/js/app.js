async function loadSongs() {
    try {
        const response = await fetch('/songs');
        const songs = await response.json();

        const songsList = document.getElementById('songs-list');
        songsList.innerHTML = '';

        songs.forEach(song => {
            const songElement = document.createElement('div');
            songElement.classList.add('song');

            songElement.innerHTML = `
                <h3>${song.song_name} - ${song.group_name}</h3>
                <p><strong>Release Date:</strong> ${song.release_date}</p>
                <p><strong>Text:</strong> ${song.text}</p>
                <p><strong>Link:</strong> <a href="${song.link}" target="_blank">Watch on YouTube</a></p>
            `;

            songsList.appendChild(songElement);
        });
    } catch (error) {
        console.error('Error loading songs:', error);
    }
}

window.onload = loadSongs;
const data = [
    {
      image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
      title: 'Luxman – 68c – Tube power amplifier',
      bid: '€140',
      timeLeft: '2 days left',
    },
    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      },    {
        image: 'https://assets.catawiki.com/image/cw_lot_card/plain/assets/catawiki/assets/2024/3/25/7/2/e/72ecfedd-7eb9-461f-a810-c268b85b8230.jpg',
        title: 'Luxman – 68c – Tube power amplifier',
        bid: '€140',
        timeLeft: '2 days left',
      }
  ];
  
  const container = document.getElementById('cardContainer');
  
  data.forEach(({ image, title, bid, timeLeft }) => {
    const card = document.createElement('div');
    card.className = 'photo-card';
  
    card.innerHTML = `
      <img src="${image}" alt="${title}" class="photo-img">
      <div class="photo-info">
        <h3 class="photo-title">${title}</h3>
        <div class="photo-details">
          <span class="label">Current bid</span>
          <span class="bid">${bid}</span>
          <span class="time">${timeLeft}</span>
        </div>
      </div>
    `;
  
    container.appendChild(card);
  });
  
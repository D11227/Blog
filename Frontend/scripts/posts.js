setTimeout(() => $('#search-input')[0].addEventListener("input", (e) => searchItems()), 1000);

let searchItem;

function searchItems() {
	if (app.currentScreen != 'home') return;
	const searchInput = $('#search-input')[0];
	searchItem = filterItems(searchInput.value);
	let html = '';
	$("#list-posts")[0].innerHTML = "";
	searchItem.forEach((post) => {
           html = '';
		   htmlTags = '';
		   post.tags.split(' ').map((x) => htmlTags += `<p class="tag">${x}</p>`);
		   html += `
	   	<div class="post" onclick="app.readPostBtn('${post.id}')">
	   		<div class="tags" style="display:flex">
	   			${htmlTags}
	   		</div>
	   		<div class="date">
	   			<p>${post.date.replace('-', '/').split('T')[0].replace('-', '/')}</p>
	   		</div>
	   		<div class="post-body">
	   			<h4>${post.title}</h4>
	   		</div>
	   	</div>
	   	`;

	$("#list-posts")[0].innerHTML += html;
	return html;
	});
}

function filterItems(post) {
         return app.posts.filter(function(item) {
            return item.title.toLowerCase().indexOf(post.toLowerCase()) > -1;
     });
}

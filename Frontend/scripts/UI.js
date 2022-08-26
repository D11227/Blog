let app;

window.onload = function() {

        app = new Vue({
                el: '#app',
                data: {
                        currentScreen: 'home',
                        server: 'http://localhost:8080/api',
                        email: '',
                        password: '',
                        posts: [],
                        content: ''
                },
                methods: {
                        getPosts: function() {
                                Fetch.get(`${this.server}/getPost`)
                                        .then(data => {
                                                if (data == []) return;
                                                data.map(x => this.posts.push(x));
                                        })
                                        .then(() => this.saveHTML());
                        },
                        saveHTML: function() {
                                let html;
                                this.posts.forEach((post) => {
                                        html = '';
                                        htmlTags = '';
                                        post.Tags.map((x) => htmlTags += `<p class="tag">${x}</p>`);
                                        html += `
					<div class="post" onclick="app.readPostBtn('${post.Id}')">
						<div class="tags" style="display:flex">
							${htmlTags}
						</div>
						<div class="date">
							<p>${post.Date.replace('-', '/').split('T')[0].replace('-', '/')}</p>
						</div>
						<div class="post-body">
							<h4>${post.Title}</h4>
						</div>
					</div>
					`;

                                        $("#list-posts")[0].innerHTML += html;
                                });
                        },
                        readPostBtn: function(id) {
                                this.currentScreen = 'post';
                                let post = this.posts.filter(x => x.Id === id)[0];

                                let html = '';
                                html += `
            	<!-- Page Content -->
            	  <div class="container" v-if="currentScreen === 'post'">
            		<div class="row">
            		  <!-- Post Content Column -->
            		  <div class="col-lg-8">
            			<!-- Title -->
            			<h1 class="mt-4">${post.Title}</h1>
            			<!-- writer -->
            			<p class="lead">
            			  by
            			  <a href="#" class="writer">${post.Writer}</a>
            			</p>
            			<hr>
            			<!-- Date/Time -->
            			<p class="date-time">Posted on ${post.Date.replace('-', '/').split('T')[0].replace('-', '/')}</p>
            			<hr>
            			<!-- Preview img -->
            			<img class="img-fluid rounded" src="${post.Image}" alt="">
            			<hr>
            			<!-- Post Content -->
            		<div class="content">
            			<div style="border: none; outline: none; height: 100%; width: 100%; overflow:auto" id="content">${post.Content}</div>
            		</div>
            			<hr>
            		<button class="btn btn-dark" style="margin: 10px;" onclick="app.back('home')">Back</button>
            		  </div>
            	</div>
            	</div>
            `
                                $('#post-blog')[0].innerHTML = html;
                                $('#post-blog').css('display', 'block');
                                setHeight('content');
                        },
                        back: function(screen) {
                                loadPage();
                                this.currentScreen = screen;
                                $('#post-blog').empty();
                                $('#post-blog').css('display', 'none');
                                if (screen == 'home') {
                                        setTimeout(this.saveHTML(), 1000);
                                }
                        },
                        isAdmin: function() {
                                fetch(`${this.server}/auth/login`, {
                                                body: new URLSearchParams(`email=${this.email}&password=${this.password}`),
                                                method: 'post'
                                        })
                                        .then(data => data.json())
                                        .then(data => document.getElementById('app').innerHTML = data.html)
                                        .then(() => {
                                                tinymce.init({
                                                        selector: 'textarea',
                                                        plugins: '3v7o6s9lm801tflvzfurcj05sjkb0l69bkxh9stldsl1bj3q',
                                                        toolbar: '3v7o6s9lm801tflvzfurcj05sjkb0l69bkxh9stldsl1bj3q',
                                                        toolbar_mode: 'floating',
                                                        tinycomments_mode: 'embedded',
                                                        tinycomments_author: 'Silverboss'
                                                });
                                        });
                        },
                        uploadPost: function() {
                                let img = document.getElementById('img').value;
                                let title = document.getElementById('title').value;
                                let tags = document.getElementById('tags').value;
                                let writer = document.getElementById('writer').value;
                                let content = tinymce.get("content").getContent();

                                const formData = {
                                        img: img,
                                        title: title,
                                        tags: tags,
                                        writer: writer,
                                        content: content
                                }

                                fetch('http://localhost:3000/post/upload', {
                                        body: JSON.stringify(formData),
                                        headers: {
                                                "Content-Type": "application/json"
                                        },
                                        method: 'post'
                                });
                        }
                }
        });
        loadTypeWritter();
        app.getPosts();
}

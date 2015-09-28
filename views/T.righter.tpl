{{define "righter"}}	
	<div class="navigation">
               

				</section>

				<aside class="prefix_1 grid_5">
					
					
					<!-- search -->
                    <form action="#" method="get">
						<input id="search" name="s" value="I'm searching for..." type="text">
					 </form>

					<!-- Categories-->
					<h4>Categories</h4>
					<ul class="categories">
						  {{range .Categories}}
							 <li><a href="/?cate={{.Title}}" >{{.Title}}</a></li>
						  {{end}}
					</ul>
					<div class="clear"></div>
					<hr>

					<!-- Comments -->
					<h4>Recent Comments</h4>
					<ul class="recent_comments">
						<li>Perfekt style here too<br><a href="#">Timo</a></li>
					</ul>
					<div class="clear"></div>
					<hr>


					<!-- Blogroll -->
					<h4>Blogroll</h4>
					<ul class="blogroll">
						<li><a href="#">Documentation</a></li>
						<li><a href="#">Plugins</a></li>
						
                    <hr>

                </aside>


            </div>
{{end}}
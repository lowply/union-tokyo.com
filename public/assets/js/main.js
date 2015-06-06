$(function(){

	//
	// Scroll to Top
	//

	$("div.control a").click(function(){
		$('body,html').animate({
			scrollTop: 0
		}, 300);
		return false;
	})

	$("div.about ul.headmenu li a[href='#union']").click(function(){
		$('body,html').animate({
			scrollTop: $("h2#union").offset().top
		}, 300);
		return false;
	})

	$("div.about ul.headmenu li a[href='#charity']").click(function(){
		$('body,html').animate({
			scrollTop: $("h2#charity").offset().top
		}, 300);
		return false;
	})

	$("div.report ul.headmenu li a[href='#charity']").click(function(){
		$('body,html').animate({
			scrollTop: $("h2#charity").offset().top
		}, 300);
		return false;
	})

	$("div.report ul.headmenu li a[href='#union']").click(function(){
		$('body,html').animate({
			scrollTop: $("h2#union").offset().top
		}, 300);
		return false;
	})

	//
	// COMING SOON
	//

	// $("ul.nav li.report a").click(function(){
	// 	return false;
	// })

	// $("ul.nav li.report a").hover(
	// 	function(){
	// 		$(this).text("COMING SOON");
	// 	},
	// 	function(){
	// 		$(this).text("REPORT");
	// 	}
	// )
});


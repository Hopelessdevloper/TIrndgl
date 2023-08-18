<!DOCTYPE html>
<html lang="en" xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">
<head>
    <meta charset="utf-8"> <!-- utf-8 works for most cases -->
    <meta name="viewport" content="width=device-width"> <!-- Forcing initial-scale shouldn't be necessary -->
    <meta http-equiv="X-UA-Compatible" content="IE=edge"> <!-- Use the latest (edge) version of IE rendering engine -->
    <meta name="x-apple-disable-message-reformatting">  <!-- Disable auto-scale in iOS 10 Mail entirely -->
    <title></title> <!-- The title tag shows in email notifications, like Android 4.4. -->

    <link href="https://fonts.googleapis.com/css?family=Lato:300,400,700" rel="stylesheet">

    <!-- CSS Reset : BEGIN -->
    <style>
::selection {
  background: aquamarine;
  color: black;
}
​
p::selection {
  background: deeppink;
  color: white;
}


        /* What it does: Remove spaces around the email design added by some email clients. */
        /* Beware: It can remove the padding / margin and add a background color to the compose a reply window. */
        html,
body {
    margin: 0 auto !important;
    padding: 0 !important;
    height: 100% !important;
    width: 100% !important;
    background: #f1f1f1;
}

/* What it does: Stops email clients resizing small text. */
* {
    -ms-text-size-adjust: 100%;
    -webkit-text-size-adjust: 100%;
}

/* What it does: Centers email on Android 4.4 */
div[style*="margin: 16px 0"] {
    margin: 0 !important;
}

/* What it does: Stops Outlook from adding extra spacing to tables. */
table,
td {
    mso-table-lspace: 0pt !important;
    mso-table-rspace: 0pt !important;
}

/* What it does: Fixes webkit padding issue. */
table {
    border-spacing: 0 !important;
    border-collapse: collapse !important;
    table-layout: fixed !important;
    margin: 0 auto !important;
}

/* What it does: Uses a better rendering method when resizing images in IE. */
img {
    -ms-interpolation-mode:bicubic;
}

/* What it does: Prevents Windows 10 Mail from underlining links despite inline CSS. Styles for underlined links should be inline. */
a {
    text-decoration: none;
}

/* What it does: A work-around for email clients meddling in triggered links. */
*[x-apple-data-detectors],  /* iOS */
.unstyle-auto-detected-links *,
.aBn {
    border-bottom: 0 !important;
    cursor: default !important;
    color: inherit !important;
    text-decoration: none !important;
    font-size: inherit !important;
    font-family: inherit !important;
    font-weight: inherit !important;
    line-height: inherit !important;
}

/* What it does: Prevents Gmail from displaying a download button on large, non-linked images. */
.a6S {
    display: none !important;
    opacity: 0.01 !important;
}

/* What it does: Prevents Gmail from changing the text color in conversation threads. */
.im {
    color: inherit !important;
}

/* If the above doesn't work, add a .g-img class to any image in question. */
img.g-img + div {
    display: none !important;
}

/* What it does: Removes right gutter in Gmail iOS app: https://github.com/TedGoas/Cerberus/issues/89  */
/* Create one of these media queries for each additional viewport size you'd like to fix */

/* iPhone 4, 4S, 5, 5S, 5C, and 5SE */
@media only screen and (min-device-width: 320px) and (max-device-width: 374px) {
    u ~ div .email-container {
        min-width: 320px !important;
    }
}
/* iPhone 6, 6S, 7, 8, and X */
@media only screen and (min-device-width: 375px) and (max-device-width: 413px) {
    u ~ div .email-container {
        min-width: 375px !important;
    }
}
/* iPhone 6+, 7+, and 8+ */
@media only screen and (min-device-width: 414px) {
    u ~ div .email-container {
        min-width: 414px !important;
    }
}

    </style>

    <!-- CSS Reset : END -->

    <!-- Progressive Enhancements : BEGIN -->
    <style>

	    .primary{
	background: #f85e9f;
}
.bg_white{
	background: #ffffff;
}
.bg_light{
	background: #fafafa;
}
.bg_black{
	background: #000000;
}
.bg_dark{
	background: rgba(0,0,0,.8);
}
.email-section{
	padding:2.5em;
}

/*BUTTON*/
.btn{
	padding: 5px 15px;
	display: inline-block;
}
.btn.btn-primary{
	border-radius: 5px;
	background: #f85e9f;
	color: #ffffff;
}
.btn.btn-white{
	border-radius: 5px;
	background: #ffffff;
	color: #000000;
}
.btn.btn-white-outline{
	border-radius: 5px;
	background: transparent;
	border: 1px solid #fff;
	color: #fff;
}
.btn.btn-black-outline{
	border-radius: 0px;
	background: transparent;
	border: 2px solid #000;
	color: #000;
	font-weight: 700;
}

h1,h2,h3,h4,h5,h6{
	font-family: 'Lato', sans-serif;
	color: #000000;
	margin-top: 0;
	font-weight: 400;
}

body{
	font-family: 'Lato', sans-serif;
	font-weight: 400;
	font-size: 15px;
	line-height: 1.8;
	color: rgba(0,0,0,.4);
}

a{
	color: #f85e9f;
}

table{
}
/*LOGO*/

.logo h1{
	margin: 0;
}
.logo h1 a{
	color: #000000;
	font-size: 20px;
	font-weight: 700;
	text-transform: uppercase;
	font-family: 'Lato', sans-serif;
	border: 2px solid #000;
	padding: .2em 1em;
}

.navigation{
	padding: 0;
	padding: 1em 0;
	/*background: rgba(0,0,0,1);*/
	border-top: 1px solid rgba(0,0,0,.05);
	border-bottom: 1px solid rgba(0,0,0,.05);
}
.navigation li{
	list-style: none;
	display: inline-block;;
	margin-left: 5px;
	margin-right: 5px;
	font-size: 13px;
	font-weight: 500;
	text-transform: uppercase;
	letter-spacing: 2px;
}
.navigation li a{
	color: rgba(0,0,0,1);
}

/*HERO*/
.hero{
	position: relative;
	z-index: 0;
}

.hero .text{
	color: rgba(0,0,0,.3);
}
.hero .text h2{
	color: #000;
	font-size: 30px;
	margin-bottom: 0;
	font-weight: 300;
}
.hero .text h2 span{
	font-weight: 600;
	color: #f85e9f;
}


/*HEADING SECTION*/
.heading-section{
}
.heading-section h2{
	color: #000000;
	font-size: 28px;
	margin-top: 0;
	line-height: 1.4;
	font-weight: 400;
}
.heading-section .subheading{
	margin-bottom: 20px !important;
	display: inline-block;
	font-size: 13px;
	text-transform: uppercase;
	letter-spacing: 2px;
	color: rgba(0,0,0,.4);
	position: relative;
}
.heading-section .subheading::after{
	position: absolute;
	left: 0;
	right: 0;
	bottom: -10px;
	content: '';
	width: 100%;
	height: 2px;
	background: #f85e9f;
	margin: 0 auto;
}

.heading-section-white{
	color: rgba(255,255,255,.8);
}
.heading-section-white h2{
	font-family: 
	line-height: 1;
	padding-bottom: 0;
}
.heading-section-white h2{
	color: #ffffff;
}
.heading-section-white .subheading{
	margin-bottom: 0;
	display: inline-block;
	font-size: 13px;
	text-transform: uppercase;
	letter-spacing: 2px;
	color: rgba(255,255,255,.4);
}


ul.social{
	padding: 0;
}
ul.social li{
	display: inline-block;
	margin-right: 10px;
}

/*FOOTER*/

.footer{
	border-top: 1px solid rgba(0,0,0,.05);
	color: rgba(0,0,0,.5);
}
.footer .heading{
	color: #000;
	font-size: 20px;
}
.footer ul{
	margin: 0;
	padding: 0;
}
.footer ul li{
	list-style: none;
	margin-bottom: 10px;
}
.footer ul li a{
	color: rgba(0,0,0,1);
}


@media screen and (max-width: 500px) {


}


    </style>


</head>

<body width="100%" style="margin: 0; padding: 0 !important; mso-line-height-rule: exactly; background-color: #222222;">
	<center style="width: 100%; background-color: #f1f1f1;">
    <div style="display: none; font-size: 1px;max-height: 0px; max-width: 0px; opacity: 0; overflow: hidden; mso-hide: all; font-family: sans-serif;">
      &zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;
    </div>
    <div style="max-width: 600px; margin: 0 auto;" class="email-container">
    	<!-- BEGIN BODY -->
      <table align="center" role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="margin: auto;">
      	<tr>
          <td valign="top" class="bg_white" style="padding: 1em 2.5em 0 2.5em;">
          	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
          		<tr>
          			<td class="logo" style="text-align: center;">
			            <h1><a href="#">Jello</a></h1>
			          </td>
          		</tr>
          	</table>
          </td>
	      </tr><!-- end tr -->
	      <tr>
          <td valign="top" class="bg_white" style="padding: 0;">
          	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
          		<tr>
			          <td width="60%" class="logo" style="text-align: center;">
			            <ul class="navigation">
			            	<li><a href="#">Contract thingy</a></li>
			            	
			            </ul>
			          </td>
          		</tr>
          	</table>
          </td>
	      </tr><!-- end tr -->
				<tr>
          <td valign="middle" class="hero hero-2 bg_white" style="padding: 2em 0 4em 0;">
            <table>
            	<tr>
            		<td>
            			<div class="text" style="padding: 0 2.5em; text-align: center;">
            				<h4>You're the <b>bestttt and deserve tons of happiness and good things</b>  </b> </h4>
            			</div>
            			<p>Also You're indebted for reasons unknown</p>
            		</td>
            	</tr>
            </table>
          </td>
	      </tr><!-- end tr -->
	      <tr>
		      <td class="bg_white">
		        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
		            		<tr>
                      
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td  class="text-services" style="text-align:  left; padding: 20px 30px;">
                            	<div class="heading-section">
								              	<h2 style="font-size: 22px;">Thanks for being you </h2>
 <p><i>In all honesty, I'll probably have to invent another phrase like "Thank you" to properly express my gratitude</i></p>         
 							              	<p>Still, I wanted to Thank you, to thank you for being the bestest most amazing  and creative, intellectual but humble person I ever met</p>
								              	<p>I particularly got inspired by the part of you that does everything she can, sometimes awkwardly and works hard to progress forward, It reminds me that I have to work just as hard myself</p>
								              								              	
								              	<p>And Thank you you for  listening to all of my problems no matter how small or insignificant they were,  for being there,  for being someone i could have conversations about random  or common things for hours, to a point that now it feels like that a part of me is gone, weird right? oh and....To be honest, you are probably the best thing this world has given me, So Thank you most of all for being you

And thank you for listening to all of my problems, no matter how small or insignificant they were, for being there, for being someone with whom I could have conversations about random or common things for hours sometimes , to the point where it now feels like a part of me is gone (weird, right?) and.... To be honest, you are probably the best thing this world has given me, so thank you most of all for being you.

</p>

<p><i>God, that was so cringy</i></p>

								              	<p></p>
								            	</div>
                            </td>
                          </tr>
                        </table>
                      </td>
                      
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
                              <img src="https://images.unsplash.com/photo-1444080748397-f442aa95c3e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTZ8fG5pZ2h0JTIwc2t5fGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=700&q=60" alt="" style="width: 100%; max-width: 680px; height: auto; margin: auto; display: block;">
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
		            		<tr>
		            			<td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
 
                            </td>
                          </tr>
                        </table>
                      </td>
                           <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td class="text-services" style="text-align: left; padding: 20px 30px;">
                            	<div class="heading-section">
								              	<h4 style="font-size: 22px;">Just so you know, For the record</h4>
								              	<p>As I wrote in the first line, you probably won't read this, but if you do,</p> 
								              	
								              	<p>I'd like you to know that I won't annoy you by sending another email asking for the reason you'd stop talking to me as you'd probably have you're own considerations backing that up.</p>

<p> Because everything that's important to begin with, the memories and friendships people build, the lives they live, everything... will be gone at some point. Everyone will take different paths. So it's okay and it'll just take me some time to get used to it.</p>
								              <p><b>And You  deserve what's best for you, Don't settle for anything less! Instead of thinking about others you should think about yourself for once! </b></p>
								            	</div>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            
      
    
                            <td class="text-services" style="text-align: left; padding: 20px 30px;">
                            	<div class="heading-section">        
              
                            	    <img src="https://images.unsplash.com/photo-1464618663641-bbdd760ae84a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8cGVhY2VmdWx8ZW58MHx8MHx8&auto=format&fit=crop&w=700&q=60" alt="" style="width: 100%; max-width: 600px; height: 200px; margin: auto; display: block;">
                            	    <h2> </h2>
								              	<h2 style="font-size: 22px;">Worries</h2>
								              	
								              	<p> 
								              I'm worried sick, even though I tried not to be. I'm honestly really worried. It's okay If you don't want to talk to me, but just in case, I wanted to ask you if something was wrong or not, or if it was just unnecessary for me to ask because I'm really really worried.
								               </p>
								              	 
								            	</div>
                            </td>
                          </tr>
                        </table>
                      </td>
                      
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
                                <h2> </h2>
                              <img src="https://images.unsplash.com/18/trickle.JPG?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTF8fHJhaW4lMjBvbiUyMHdpbmRvd3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=700&q=60" alt="" style="width: 100%; max-width: 600px; height: 200px; margin: auto; display: block;">
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
		            		<tr>
		            			<td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
                              <img src="images/product-4.jpg" alt="" style="width: 100%; max-width: 600px; height: auto; margin: auto; display: block;">
                            </td>
                          </tr>
                        </table>
                      </td>
                      <td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td class="text-services" style="text-align: left; padding: 20px 30px;">
                            	<div class="heading-section">
								              	<h2 style="font-size: 22px;">Take care of yourself</h2>
								              	<p>You know already, but please don't sacrifice yourself too much you deserve better, you know you do. Your mental health is just as important as your physical one!</p>
								             
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
		            		<tr>
                      <td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td class="text-services" style="text-align: left; padding: 20px 30px;">
                            	<div class="heading-section">
								              	
								              	<
                            </td>
                          </tr>
                        </table>
                      </td>
                      <td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
                              <img src="images/product-5.jpg" alt="" style="width: 100%; max-width: 600px; height: auto; margin: auto; display: block;">
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		          <tr>
		            <td class="bg_light email-section" style="padding: 0; width: 100%;">
		            	<table role="presentation" border="0" cellpadding="0" cellspacing="0" width="100%">
		            		<tr>
		            			<td valign="middle" width="50%">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td>
                              <img src="images/product-6.jpg" alt="" style="width: 100%; max-width: 600px; height: auto; margin: auto; display: block;">
                            </td>
                          </tr>
                        </table>
                      </td>
                    
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                          <tr>
                            <td class="text-services" style="text-align: left; padding: 20px 30px;">
                            	
								              	<p><a href="#" class="btn btn-primary">Sincerely -</a></p>
								            	</div>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
		            	</table>
		            </td>
		          </tr><!-- end: tr -->
		        </table>
		      </td>
		    </tr><!-- end:tr -->
		    <tr>
          <td valign="middle" class="bg_white" style="padding: 2em 0;">
            <table>
            	<tr>
            		<td>
            			<div class="text" style="padding: 0 2.5em; text-align: center;">
            				<p><a href="#" class="btn btn-black-outline">Jelly</a></p>
            			</div>
            		</td>
            	</tr>
            </table>
          </td>
	      </tr><!-- end tr -->
      <!-- 1 Column Text + Button : END -->
      </table>
      <table align="center" role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="margin: auto;">
      	<tr>
          <td valign="middle" class="bg_white footer">
        
                
         
                		
                		</li>
                	</ul>
                </td>
              </tr>
            </table>
          </td>
        </tr><!-- end: tr -->
        <tr>
          <td class="bg_light" style="text-align: center;">
         <p>Worried about spam? <a href="#" style="color: rgba(0,0,0,.8);">No worries no further emails will be sent</a></p>
          </td>
        </tr>
      </table>

    </div>
  </center>
</body>
</html>

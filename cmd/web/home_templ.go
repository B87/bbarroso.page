// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white shadow-md rounded p-6 mx-8\"><div class=\"flex flex-col md:flex-row\"><div class=\"w-full md:w-1/4 flex justify-center items-center mb-4 md:mb-0\"><img src=\"/assets/images/profilepic.jpg\" alt=\"Profile Picture\" class=\"rounded-full w-40 h-40 object-cover\"></div><div class=\"w-full md:w-3/4 md:pl-4 md:mr-10\"><h1 class=\"text-2xl font-bold mb-2 mt-4 md:mt-0\">Bernat Barroso</h1><div class=\"mt-1\"><div class=\"flex flex-wrap space-x-4\"><a href=\"https://www.linkedin.com/in/bernat-barroso-81210a94/\" class=\"text-blue-900\">LinkedIn</a> <a href=\"https://www.github.com/B87\" class=\"text-blue-900\">GitHub</a></div></div><p class=\"mt-2 text-lg mr-10\">I focus on solving business problems by applying process engineering and automation.<br>Throughout my career, I've learned crucial IT and management skills for the current data centric landscape, applying them to resolve key business operations challenges and consistently deliver impactful results.</p></div></div><div class=\"mt-4 p-5 mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Work Experience</h2><ul class=\"relative pl-10 mt-5\"><li class=\"relative mb-7 pl-5 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-0.5 before:bg-teal-600\"><p class=\"font-bold\">GINDUMAC - CTO</p><p class=\"italic text-gray-700 mb-2.5\">September 2019 - Present</p><p>Responsible for all technology assets of the company.<br>The platform is based on AWS services and multiple programming languages, mainly PHP, Javascript and Python.</p></li><li class=\"relative mb-7 pl-5 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-0.5 before:bg-teal-600\"><p class=\"font-bold\">GINDUMAC - Data Engineer</p><p class=\"italic text-gray-700 mb-2.5\">October 2018 - July 2019</p><p>Architect and developer of MIS, an application which provides relevant market data scrapped from multiple websites to internal personnel in order to optimize their decisions and operations. It is based on AWS services and written in Python and Java.</p></li><li class=\"relative mb-7 pl-5 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-0.5 before:bg-teal-600\"><p class=\"font-bold\">Everis - Data Engineer</p><p class=\"italic text-gray-700 mb-2.5\">July 2016 - November 2018</p><p>Owner of diverse Apache Spark data processing projects for a financial institution, including the technical documentation, testing and performance analysis. All of them were based on an on-premises CDH Hadoop cluster and written in Java.</p></li><li class=\"relative mb-7 pl-5 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-0.5 before:bg-teal-600\"><p class=\"font-bold\">Everis - Software Architecture Intern</p><p class=\"italic text-gray-700 mb-2.5\">October 2015 - June 2016</p><p>During this internship I discovered my passion for software, building ETL pipelines using SQL and Java in a big data environment.</p></li><li class=\"relative mb-7 pl-5 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-0.5 before:bg-teal-600\"><p class=\"font-bold\">Accenture - Supply Chain Management Intern</p><p class=\"italic text-gray-700 mb-2.5\">July 2015 - September 2015</p><p>Worked there helping the analytics team, validating and cleaning data. We provided forecasts for supply chain optimization.</p></li></ul></div><div class=\"mt-4 education-section p-5 rounded-lg  mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Formal Education</h2><ul class=\"list-none p-0\"><li class=\"mb-5\"><p class=\"font-bold mb-1\">Universitat Autònoma de Barcelona (UAB) - B.A in Business and Information Technology</p><p class=\"italic text-gray-600 mb-2\">Graduated: July 2016</p><p class=\"leading-relaxed text-gray-800\">This bachelor's degree aims to provide the alumni with traditional business knowledge and skills, such as marketing, accounting, or macroeconomics, plus knowledge of how the use of information technology enables the improvement of existing business models and disruptive new ones.</p></li></ul></div><div class=\"mt-4 education-section p-5 rounded-lg  mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Formal Education</h2><ul class=\"education-list list-none p-0\"><li class=\"education-item mb-5\"><p class=\"font-bold mb-1\">Universitat Autònoma de Barcelona (UAB) - B.A in Business and Information Technology</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: July 2016</p><p class=\"education-description leading-relaxed text-gray-800\">This bachelor's degree aims to provide the alumni with traditional business knowledge and skills, such as marketing, accounting, or macroeconomics, plus knowledge of how the use of information technology enables the improvement of existing business models and disruptive new ones.</p></li></ul></div><div class=\"mt-4 p-5 mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Certifications</h2><ul class=\"certification-list list-none p-0\"><li class=\"certification-item mb-5\"><p class=\"font-bold mb-1\">Convolutional Neural Networks - Coursera</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: December 2017</p><p><a href=\"https://www.coursera.org/account/accomplishments/verify/BTRCKT4PQM2L\" class=\"certificate-link text-blue-900 underline\">Certificate Link</a></p></li><li class=\"certification-item mb-5\"><p class=\"font-bold mb-1\">Improving Deep Neural Networks: Hyperparameter tuning, Regularization and Optimization - Coursera</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: October 2017</p><p><a href=\"https://www.coursera.org/account/accomplishments/verify/NKP26TZLWS9Y\" class=\"certificate-link text-blue-900 underline\">Certificate Link</a></p></li><li class=\"certification-item mb-5\"><p class=\"font-bold mb-1\">Neural Networks and Deep Learning - Coursera</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: October 2017</p><p><a href=\"https://www.coursera.org/account/accomplishments/verify/C9EB9QKZ4RDQ\" class=\"certificate-link text-blue-900 underline\">Certificate Link</a></p></li><li class=\"certification-item mb-5\"><p class=\"font-bold mb-1\">Machine Learning - Coursera</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: April 2016</p><p><a href=\"https://www.coursera.org/account/accomplishments/verify/MGNRV6CGMAUV\" class=\"certificate-link text-blue-900 underline\">Certificate Link</a></p></li><li class=\"certification-item mb-0\"><p class=\"font-bold mb-1\">M101J : MongoDB for Java Developers - MongoDB, Inc.</p><p class=\"graduation-date italic text-gray-600 mb-2\">Graduated: February 2016</p><p><a href=\"http://university.mongodb.com/course_completion/82bb563c45924d8f82c6601cf4d34ada\" class=\"certificate-link text-blue-900 underline\">Certificate Link</a></p></li></ul></div><div class=\"mt-4 p-5 mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Skills</h2><div class=\"flex flex-col md:flex-row flex-wrap space-y-4 md:space-y-0 md:space-x-4\"><div class=\"bg-gray-200 rounded px-4 py-2\">Golang</div><div class=\"bg-gray-200 rounded px-4 py-2\">Python</div><div class=\"bg-gray-200 rounded px-4 py-2\">GCP / AWS</div><div class=\"bg-gray-200 rounded px-4 py-2\">AI</div><div class=\"bg-gray-200 rounded px-4 py-2\">PostgreSQL</div><div class=\"bg-gray-200 rounded px-4 py-2\">Elasticsearch</div><div class=\"bg-gray-200 rounded px-4 py-2\">MongoDB</div></div></div><div class=\"mt-4 p-5 mx-8\"><h2 class=\"text-xl font-semibold mb-5 text-teal-600\">Testimonials</h2><ul class=\"list-none p-0\"><li class=\"mb-5\"><p class=\"italic text-gray-800 mb-2\">\"Working with Bernat was a very interesting experience. It is not often that you have somebody who is able to grasp a problem with all its details, and at the same time think outside of the box and be pragmatic. His way of solving problems in the design of the data infrastructure was very creative and helped to reduce the complexity by a very high margin. Very highly recommended.\"</p><p class=\"font-semibold text-gray-600\">- Daniel Sapundziev - GINDUMAC Head of IoT & Data Science</p></li></ul></div></div><style>\n\t\t\t/* Timeline Styles */\n\t\t\t.timeline {\n\t\t\t\tposition: relative;\n\t\t\t\tpadding-left: 40px;\n\t\t\t\tmargin-top: 20px;\n\t\t\t}\n\n\t\t\t.timeline::before {\n\t\t\t\tcontent: '';\n\t\t\t\tposition: absolute;\n\t\t\t\ttop: 0;\n\t\t\t\tbottom: 0;\n\t\t\t\tleft: 20px;\n\t\t\t\twidth: 2px;\n\t\t\t\tbackground: #00796b;\n\t\t\t}\n\n\t\t\t.timeline-item {\n\t\t\t\tposition: relative;\n\t\t\t\tmargin-bottom: 30px;\n\t\t\t\tpadding-left: 20px;\n\t\t\t}\n\n\t\t\t.timeline-item::before {\n\t\t\t\tcontent: '';\n\t\t\t\tposition: absolute;\n\t\t\t\tleft: -11px;\n\t\t\t\ttop: 0;\n\t\t\t\twidth: 10px;\n\t\t\t\theight: 10px;\n\t\t\t\tborder-radius: 50%;\n\t\t\t\tbackground: #00796b;\n\t\t\t\tborder: 2px solid #fff;\n\t\t\t}\n\n\t\t\t.timeline-item p {\n\t\t\t\tmargin: 5px 0;\n\t\t\t}\n\n\t\t\t.timeline-item .years {\n\t\t\t\tfont-style: italic;\n\t\t\t\tcolor: #555;\n\t\t\t\tmargin-bottom: 10px;\n\t\t\t}\n\n\t\t\t.timeline-item:last-child {\n\t\t\t\tmargin-bottom: 0;\n\t\t\t}\n\n\n\t\t</style>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

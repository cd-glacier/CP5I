//
//  ViewController2.swift
//  CP5I
//
//  Created by glacier on 2017/09/15.
//  Copyright © 2017年 Bteam. All rights reserved.
//

import UIKit
import Alamofire
import SwiftyJSON
import AlamofireImage

class ViewController2: UIViewController {

    @IBOutlet weak var recipeNameLabel: UILabel!
    @IBOutlet weak var recipeImageView: UIImageView!
    @IBOutlet weak var ingredientTextView: UITextView!
    @IBOutlet weak var methodTextView: UITextView!
    
	//var ingredients: [Ingredient]
    
    var id = 1

	override func viewDidLoad() {
		super.viewDidLoad()
        //ingredientTextView.isEditable = false
        //methodTextView.isEditable = false

		// Do any additional setup after loading the view.

		let url:String = "http://noticeweb.net/api/recipe/" + String(id)
		Alamofire.request(url).responseJSON { response in
			debugPrint(url)
			let json = JSON(response.result.value)
			let data = json["data"]
            self.recipeNameLabel.text = data["name"].string!
            self.recipeImageView.af_setImage(withURL: NSURL(string: data["image_url"].string!)! as URL, placeholderImage: UIImage(named: "hoiru.png"), imageTransition: .crossDissolve(0.5))
            data["ingredients"].forEach{(_, ingredient) in
                var name:String = ingredient["name"].string!
                self.ingredientTextView.text = self.ingredientTextView.text + String(name) + "\n"
            }
            data["method"].forEach{(i, m) in
                var content:String = m["content"].string!
                self.methodTextView.text = self.methodTextView.text + String(i) + ". " + String(content) + "\n"
            }
            
		}
	}
	
	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
		// Dispose of any resources that can be recreated.
	}

	/*
	// MARK: - Navigation

	// In a storyboard-based application, you will often want to do a little preparation before navigation
	override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
	// Get the new view controller using segue.destinationViewController.
	// Pass the selected object to the new view controller.
	}
	*/
}


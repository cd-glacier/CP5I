//
//  ViewController.swift
//  CP5I
//
//  Created by glacier on 2017/09/12.
//  Copyright © 2017年 Bteam. All rights reserved.
//

import UIKit
import Alamofire
import SwiftyJSON

class ViewController: UIViewController {

    
    @IBOutlet weak var textView: UITextView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        
        Alamofire.request("http://noticeweb.net/api/easy/recipe").responseJSON { response in
            let json = JSON(response.result.value)
            json["data"].forEach{(_, data) in
                let type = data["name"].string!
                self.textView.text = String(type)
            }
        }
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

